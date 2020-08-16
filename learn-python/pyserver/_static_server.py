import socket


"""
    返回固定页面的简单http服务器
"""


def service_client(new_socket):
    req = new_socket.recv(1024)
    print("request: {req}".format(req=req))

    resp = "HTTP/1.1 200 OK\r\n"
    resp += "\r\n"
    resp += "<h1>Hello Python</h1>"
    new_socket.send(resp.encode("utf8"))

    new_socket.close()


def main():
    tcp_server_socket = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    tcp_server_socket.bind(("127.0.0.1", 8080))
    tcp_server_socket.listen(128)

    while True:
        new_socket, client_addr = tcp_server_socket.accept()
        print("client addr: {}".format(client_addr))
        service_client(new_socket)

    tcp_server_socket.close()


if __name__ == '__main__':
    main()