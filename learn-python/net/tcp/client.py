import socket


def client():
    # 客户端套接字
    tcp_client_socket = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    # 尝试连接服务端
    tcp_client_socket.connect(('127.0.0.1', 8000))

    while True:
        send_msg = input('向服务端发送消息：')
        if send_msg:
            tcp_client_socket.send(send_msg.encode('utf8'))
        else:
            break

        server_msg = tcp_client_socket.recv(1024)
        print('服务端应答：{msg}'.format(msg=server_msg.decode('utf8')))

    tcp_client_socket.close()


if __name__ == '__main__':
    client()