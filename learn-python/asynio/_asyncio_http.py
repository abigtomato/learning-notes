import asyncio
from urllib.parse import urlparse
import time


"""
    asyncio模拟http请求
"""


async def get_url(url):
    url = urlparse(url)
    host = url.netloc
    path = url.path
    if path == "":
        path = "/"
    
    # 建立socket是阻塞操作，使用await挂起函数切换到其他协程执行
    # open_connection是一个协程，用于建立socket连接
    # open_connection返回接收对象reader和发送对象writer
    reader, writer = await asyncio.open_connection(host, 80)
    # 发送数据，注销写事件，注册读事件
    writer.write("GET {} HTTP/1.1\r\nHost:{}\r\nConnection:close\r\n\r\n".format(path, host).encode("utf8"))
    
    # async for语法，异步接收数据(每次执行recv(1024)操作)
    all_lines = []
    async for raw_line in reader:
        data = raw_line.decode("utf8")
        all_lines.append(data)

    # join接收到的所有内容
    html = "\n".join(all_lines)
    return html


async def main():
    tasks = []
    for url in range(20):
        url = "http://www.baidu.com"
        # 每次循环都新注册一个get_url协程
        tasks.append(asyncio.ensure_future(get_url(url)))
    
    # as_completed()的方式可以批量执行多个协程任务并获取结果
    for task in asyncio.as_completed(tasks):
        # await暂停任务等待协程执行
        result = await task
        print(result)


if __name__ == '__main__':
    start_time = time.time()
    
    # 事件循环
    loop = asyncio.get_event_loop()
    # 注册main协程
    loop.run_until_complete(main())

    print("last time: {time}".format(time=time.time() - start_time))