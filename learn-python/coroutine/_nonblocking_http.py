import socket
from urllib.parse import urlparse


"""
    socket非阻塞io模拟http请求
"""


def get_url(url): 
    url = urlparse(url=url)
    host = url.netloc
    path = url.path
    if path == '':
        path = '/'

    client = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    # 设置为非阻塞式io
    client.setblocking(False)

    try:
        # 非阻塞式io不会阻塞等待连接建立，而是直接返回，此时未建立成功会抛出异常，所以添加捕获代码块
        client.connect((host, 80))
    except BlockingIOError:
        pass

    while True:
        try:
            # 循环不断的尝试发送请求，直到tcp连接建立，发送成功后退出循环
            client.send('GET {path} HTTP/1.1\r\nHost:{host}\r\nConnection:close\r\n\r\n'
                        .format(path=path, host=host).encode('utf-8'))
            break
        except OSError:
            pass

    data = b''
    while True:
        try:
            # 循环不断的尝试接收响应，若抛出异常则是本次未接收到响应数据，这时结束本次循环，执行下一次循环尝试接收响应
            c_data = client.recv(1024)
        except BlockingIOError:
            continue
        if c_data:
            data += c_data
        else:
            break

    data = data.decode('utf8')
    splits = data.split('\r\n\r\n')

    print(splits[0])
    print(splits[1])

    client.close()


if __name__ == '__main__':
    get_url("http://www.baidu.com")