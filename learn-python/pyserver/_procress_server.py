import socket
import re
import multiprocessing


"""
   多进程并发服务器
"""


def service_client(new_socket):
    req = new_socket.recv(1024)
    req_list = req.decode("utf8").splitlines()
    print("request: {req}".format(req=req_list))

    # GET /index.html HTTP/1.1 正则抽取文件名
    ret = re.match(r"[^/]+/([^ ]*)", req_list[0])
    if ret:
        file_name = ret.group(1)
        # 若请求结果不存在文件路径，则默认访问主页
        if file_name == "":
            file_name = "index.html"

    try:
        f = open("E:\\usr\\learn-python\\03\\http server\\html\\" + file_name, "rb")
    except Exception as e:
        # 存在异常情况
        print("error info: {info}".format(info=e))
        resp = "HTTP/1.1 404 NOT FOUND\r\n"
        resp += "\r\n"
        resp += "<h1>404</h1>"
        new_socket.send(resp.encode("utf8"))
    else:
        # 未捕获到异常情况
        html_content = f.read()
        f.close()

        resp = "HTTP/1.1 200 OK\r\n"
        resp += "\r\n"
        new_socket.send(resp.encode("utf8"))
        new_socket.send(html_content)

    new_socket.close()


def main():
    tcp_server_socket = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    tcp_server_socket.setsockopt(socket.SOL_SOCKET, socket.SO_REUSEADDR, 1)
    
    tcp_server_socket.bind(("127.0.0.1", 8080))
    tcp_server_socket.listen(128)

    while True:
        new_socket, client_addr = tcp_server_socket.accept()
        print("client addr: {addr}".format(addr=client_addr))
        
        process = multiprocessing.Process(target=service_client, args=(new_socket, ))
        process.start()

        # socket在linux中是以文件的形式存在，通过文件描述符区分
        # 多进程写时拷贝的特性会将socket的文件描述符拷贝一份，相当于此文件存在多个软链接
        # 在service_client中close掉socket时只是减少一层软链接而文件不受影响，所有在外部再次调用close关闭socket
        new_socket.close()        

    tcp_server_socket.close()


if __name__ == '__main__':
    main()