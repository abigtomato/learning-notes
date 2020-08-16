import socket


def server():
    # AF_INET ipv4形式
    # SOCK_STREAM tcp协议
    tcp_server_socket = socket.socket(socket.AF_INET, socket.SOCK_STREAM)

    # 服务端绑定ip地址和端口
    tcp_server_socket.bind(('127.0.0.1', 8000))
    # 使套接字被动接受链接
    tcp_server_socket.listen(128)
    
    # 循环不断的为连接上的客户端服务
    while True:
        # 阻塞式监听，等待客户端发送的连接请求
        # new_client_socket 新连接上的客户端套接字
        # client_addr 连接上的客户端的信息
        new_client_socket, client_addr = tcp_server_socket.accept()
        print('与{client}建立连接'.format(client=client_addr))

        # 循环不断的为连接上的客户端服务
        while True:
            recv_data = new_client_socket.recv(1024)
            if recv_data:
                print('{client}说：{msg}'.format(client=client_addr, msg=recv_data.decode('utf8')))
            else:
                break

            send_data = input('向客户端回返消息：')
            new_client_socket.send('{msg}'.format(msg=send_data).encode('utf8'))

        new_client_socket.close()

    tcp_server_socket.close()


if __name__ == '__main__':
    server()
