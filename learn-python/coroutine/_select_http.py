import socket
from urllib.parse import urlparse
from selectors import DefaultSelector, EVENT_WRITE, EVENT_READ


"""
    io多路复用select模拟http请求(单线程并发)]
    select + 事件循环 + 回调
"""


# DefaultSelector()会根据操作系统选择不同的io多路复用机制
# windows只支持select，unix/linux支持select和poll/epoll
selector = DefaultSelector()
urls = []
stop = False


class Fetcher(object):
    # EVENT_WRITE写事件(网络io中的发送操作)的回调函数
    # 当有socket连接触发写事件会被事件循环函数扫描到并调用绑定的回调方法
    def connected(self, key):
        # unregister()取消当前连接的注册
        # key.fd就是取得当前socket连接文件句柄
        selector.unregister(key.fd)
        self.client.send("GET {} HTTP/1.1\r\nHost:{}\r\nConnection:close\r\n\r\n".format(self.path, self.host).encode("utf8"))
        # 为socket连接注册读事件并绑定回调函数
        selector.register(self.client.fileno(), EVENT_READ, self.readable)

    # EVENT_READ读事件(网络io中的接收操作)的回调函数
    def readable(self, key):
        d = self.client.recv(1024)
        if d:
            self.data += d
        else:
            selector.unregister(key.fd)
            data = self.data.decode("utf8")
            splits = data.split("\r\n\r\n")

            print(splits[0])
            print(splits[1])
            self.client.close()
            # 删除处理完的url
            urls.remove(self.spider_url)

            # 若所有url处理完毕，则结束事件循环
            if not urls:    
                global stop
                stop = True

    def get_url(self, url):
        self.spider_url = url
        url = urlparse(url)
        self.host = url.netloc
        self.path = url.path
        self.data = b""
        if self.path == "":
            self.path = "/"

        self.client = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
        # 设置为非阻塞式io
        self.client.setblocking(False)

        try:
            self.client.connect((self.host, 80))
        except BlockingIOError:
            pass

        # 为socket连接注册写事件并绑定回调函数
        # fileno()表示文件句柄，select会根据句柄判断不同的事件触发
        selector.register(self.client.fileno(), EVENT_WRITE, self.connected)


# 事件循环
def loop():
    # 使用标记变量控制循环体，避免无socket连接依旧循环监控事件的情况
    while not stop:
        # select事件循环机制，监控socket连接是否触发事件
        ready = selector.select()
        for key, mask in ready:
            print(mask)
            # key.data就是相应的回调函数名
            call_back = key.data 
            call_back(key)


if __name__ == '__main__':
    url = "http://www.baidu.com"
    urls.append(url)

    fetcher = Fetcher()
    fetcher.get_url(url)
    loop()
