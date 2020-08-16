import socket


def send_file_2_client(new_client_socket):
    file_name = new_client_socket.recv(1024)

    file_content = None
    try:
        f = open('./data/{name}'.format(name=file_name.decode('utf8')), 'rb')
        file_content = f.read()
        f.close()
    except Exception as e:
        print('{exception}'.format(exception=e))

    if file_content:
        new_client_socket.send(file_content)


# 文件下载服务端
def server():
    tcp_server = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    tcp_server.bind(('127.0.0.1', 8000))
    tcp_server.listen(128)

    while True:
        new_client_socket, client_addr = tcp_server.accept()
        print('与{client}建立连接'.format(client=client_addr))

        send_file_2_client(new_client_socket)

        new_client_socket.close()

    tcp_server.close()


if __name__ == '__main__':
    server()