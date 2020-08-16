import socket


# 文件下载客户端
def client():
    tcp_client = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    tcp_client.connect(('127.0.0.1', 8000))

    file_name = input('输入要下载的文件名：')
    tcp_client.send(file_name.encode('utf8'))

    recv_data = tcp_client.recv(1024)
    if recv_data:
        with open('./data/{name}'.format(name=file_name), 'wb') as f:
            f.write(recv_data)

    tcp_client.close()


if __name__ == '__main__':
    client()