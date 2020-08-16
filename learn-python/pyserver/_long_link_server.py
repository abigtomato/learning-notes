import re
import socket


"""
    基于长链接的单进程，单线程并发服务器
"""


# 负责为单个客户端服务
def service_client(new_socket, request):
    # 1.正则抽取客户端请求的资源路径
    req_lines = request.splitlines()
    # GET /index.html HTTP/1.1
    ret = re.match(r"[^/]+/([^ ]*)", req_lines[0])
    if ret:
        file_name = ret.group(1)
        if file_name == "":
            file_name = "index.html"
    
    try:
        f = open("E:\\usr\\learn-python\\03\\http server\\html\\" + file_name, "rb")
    except Exception as ret:
        # 2.若服务端无此资源，则响应404
        response = "HTTP/1.1 404 not found\r\n"
        response += "\r\n"
        response += "<h1>file not found</h1>"
        new_socket.send(response.encode("utf-8"))
    else:
        # 3.若存在资源，则拼接响应头和响应体为响应数据
        html_content = f.read()
        f.close()
        response_body = html_content

        # http1.1版本支持长连接
        response_header = "HTTP/1.1 200 ok\r\n"
        # 设置响应数据的长度，这样客户端才能判断此次接收是否完毕
        response_header += "Content-Length:{len}\r\n".format(len=len(response_body))
        response_header += "\r\n"

        response = response_header.encode("utf-8") + response_body
        # 4.发送响应数据，但不关闭客户端套接字，因为长连接会反复利用此次的连接
        new_socket.send(response)


# 负责监听和建立与客户端的连接
def main():
    # 1.创建监听套接字
    tcp_server_socket = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    tcp_server_socket.setsockopt(socket.SOL_SOCKET, socket.SO_REUSEADDR, 1)

    # 2.绑定地址端口
    tcp_server_socket.bind(("127.0.0.1", 7890))
    tcp_server_socket.listen(128)

    # 3.设置非阻塞
    tcp_server_socket.setblocking(False)
    client_socket_list = list()

    while True:
        try:
            # 4.循环等待tcp连接建立
            client_socket, client_addr = tcp_server_socket.accept()
        except Exception:
            pass
        else:
            print("与 {info} 建立连接".format(info=client_addr))
            client_socket.setblocking(False)
            client_socket_list.append(client_socket)

        for client_socket in client_socket_list:
            try:
                # 5.接收建立连接的客户端发送的消息
                recv_data = client_socket.recv(1024)
            except Exception:
                pass
            else:
                if recv_data:
                    # 6.为每一个客户端单独服务
                    service_client(client_socket, recv_data.decode("utf-8"))
                else:
                    # 7.若客户端无数据发生，则关闭客户端套接字
                    client_socket.close()
                    client_socket_list.remove(client_socket)
    else:
        # 8.若服务结束，则关闭监听套接字
        tcp_server_socket.close()


if __name__ == '__main__':
    main()