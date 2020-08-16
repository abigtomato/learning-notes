import socket
import re


"""
    动态返回页面的简单http服务器
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
        f = open("E:\\usr\\learn-python\\03\\web server\\html\\" + file_name, "rb")
    except Exception as e:
        # 存在异常情况
        print("error info: {info}".format(e))
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
        print("client addr: {info}".format(info=client_addr))
        service_client(new_socket)

    tcp_server_socket.close()


if __name__ == '__main__':
    main()