import re
import select
import socket


"""
    io多路复用epoll版单进程，单线程并发服务器
    1.epoll在内存的用户态和内核态间使用内存映射技术建立了一块公有的内存空间
    2.当socket的文件描述符注册到epoll中就相当于存入此内存空间，操作系统内核会检测此块内存中的fd是否触发绑定的事件
    3.一旦检测到事件触发文件描述符就绪，就采用类似回调的机制做出响应
"""

 
def service_client(new_socket, request):
    req_lines = request.splitlines()
    ret = re.match(r"[^/]+/([^ ]*)", req_lines[0])
    if ret:
        file_name = ret.group(1)
        if file_name == "":
            file_name = "index.html"
    
    try:
        f = open("E:\\usr\\learn-python\\03\\http server\\html\\" + file_name, "rb")
    except Exception:
        response = "HTTP/1.1 404 NOT FOUND\r\n"
        response += "\r\n"
        response += "<h1>file not found</h1>"
        new_socket.send(response.encode("utf-8"))
    else:
        html_content = f.read()
        f.close()

        response_body = html_content
        response_header = "HTTP/1.1 200 OK\r\n"
        response_header += "Content-Length:{}\r\n".format(len(response_body))
        response_header += "\r\n"

        response = response_header.encode("utf-8") + response_body
        new_socket.send(response)


def main():
    tcp_server_socket = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    tcp_server_socket.setsockopt(socket.SOL_SOCKET, socket.SO_REUSEADDR, 1)
    
    tcp_server_socket.bind(("127.0.0.1", 7890))
    tcp_server_socket.listen(128)

    tcp_server_socket.setblocking(False)
    
    # 创建epoll对象
    epl = select.epoll()
    # 将监听套接字对应的文件描述符注册到epoll中
    # select.EPOLLIN表示为文件描述符绑定输入事件(若是有数据输入，则代表有客户端建立连接，则触发事件)
    epl.register(tcp_server_socket.fileno(), select.EPOLLIN)
    
    fd_event_dict = dict()

    while True:
        # 默认阻塞，直到os检测到数据到来触发事件再解除阻塞
        fd_event_list = epl.poll()
        # (fd, event) => (fd是套接字对应的文件描述符，event是文件描述符对应的事件，如:recv接收事件)
        for fd, event in fd_event_list:
            # 若是监听套接字触发事件则代表有新的客户端套接字连接
            if fd == tcp_server_socket.fileno():
                client_socket, client_addr = tcp_server_socket.accept()
                print("与 {info} 建立连接".format(info=client_addr))

                # 将新的客户端套接字注册到epoll中
                epl.register(client_socket.fileno(), select.EPOLLIN)
                fd_event_dict[client_socket.fileno()] = client_socket
            # 若是客户端套接字触发输入事件则代表有新数据等待接收
            elif event == select.EPOLLIN:
                client_socket = fd_event_dict[fd]
                recv_data = client_socket.recv(1024)
                if recv_data:
                    service_client(client_socket, recv_data)
                else:
                    fd_event_dict[fd].close()
                    # 服务结束后将此客户端套接字从epoll中取消注册
                    epl.unregister(fd)
                    del fd_event_dict[fd]
    else:
        # 关闭监听套接字
        tcp_server_socket.close()


if __name__ == '__main__':
    main()