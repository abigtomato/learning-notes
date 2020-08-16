import asyncio
import aiohttp
import aiomysql

import re
from pyquery import PyQuery


"""
    aiohttp协程并发爬虫
"""


# 消费循环结束的标记
stopping = False
# 起始的url
start_url = "http://www.jobbole.com/"
# 待爬取的url
waitting_urls = list()
# 爬取完成等待去重的url
seen_urls = set() 

# 控制协程并发数为3
sem = asyncio.Semaphore(3)


# 发送http请求的协程
async def fetch(url, session):
    # 协程并发数为3
    async with sem:
        try:
            # 通过session会话对象发送http请求并获取响应对象resp
            async with session.get(url) as resp:
                print("{url} 响应状态: {status}".format(url=url, status=resp.status))
                if resp.status in [200, 201]:
                    html = await resp.text()
                    print("响应体内容: {html}".format(html=html))
        except Exception as e:
            print("Exception: {info}".format(info=e))


# 从html中提取url的函数
def extract_urls(html):
    urls = []
    pq = PyQuery()
    # 使用PyQuery()从html中抽取url
    for link in pq.items("a"):
        url = link.attr("href")
        if url and url.startswith("http") and url not in seen_urls:
            urls.append(url)
            # 存入等待队列
            waitting_urls.append(url)
    return urls


# 初始化url的协程
async def init_urls(url, session):
    html = await fetch(url, session)
    seen_urls.add(url)
    extract_urls(html)


# 处理文章详情页并解析入库的协程
async def article_handler(url, session, pool):
    html = await fetch(url, session)
    # 存入完成队列
    seen_urls.add(url)
    # 从html提取url并存入等待队列
    extract_urls(html)
    
    # 提取标题并入库
    pq = PyQuery(html)
    title = pq("title").text()
    async with pool.acquire() as conn:
        async with conn.cursor() as cur:
            # await cur.execute("SELECT 42;")
            insert_sql = "insert into article_test(title) values('{val}')".format(val=title)
            # 提交sql
            await cur.execute(insert_sql)


# 消费waitting_urls等待队列中的url的协程
async def consumer(pool):
    async with aiohttp.ClientSession() as session:
        # 循环消费url并解析入库
        while not stopping:
            # 若等待队列中无可消费的url，则添加耗时操作暂停消费协程
            if len(waitting_urls) == 0:
                await asyncio.sleep(0.5)
                continue

            url = waitting_urls.pop()
            print("开始获取: {url}".format(url=url))

            if re.match("http://.*?jobbole.com/\d+/", url):
                if url not in seen_urls:
                    # 解析入库
                    asyncio.ensure_future(article_handler(url, session, pool))
            else:
                if url not in seen_urls:
                    asyncio.ensure_future(init_urls(url, session))


# 调度协程
async def main(loop):
    # aiomysql.create_pool()创建异步的mysql连接
    # 在阻塞等待mysql连接池连接建立的时候使用await挂起，执行其他协程
    pool = await aiomysql.create_pool(host="127.0.0.1", port=3306,
        user="root", password="1234", db="aiomysql_test", loop=loop,
        charset="utf8", autocommit=True)

    # aiohttp.ClientSession()发送http请求的客户端会话对象
    async with aiohttp.ClientSession() as session:
        # await使当前协程暂停，调度fetch协程请求起始页的url
        html = await fetch(start_url, session)
        # 添加起始页到爬取完毕的集合中
        seen_urls.add(start_url)
        # 调用extract_urls()函数从起始页的html中提取url(此时url都存储在待爬取队列waitting_urls中)
        extract_urls(html)
    # 将consumer消费协程注册到asyncio的事件循环中
    asyncio.ensure_future(consumer(pool))


if __name__ == '__main__':
    # 获取asyncio的事件循环
    loop = asyncio.get_event_loop()
    # 注册main调度协程(asyncio可以看做协程池，注册协程后无需关心协程如何调度，由事件循环监听触发事件后调度)
    asyncio.ensure_future(main(loop))
    # 开启事件循环
    loop.run_forever()