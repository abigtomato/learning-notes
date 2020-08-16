import socket


def send_msg(udp_socket):
    send_data = input('请输入：')
    udp_socket.sendto(send_data.encode('utf8'), ('127.0.0.1', 8000))


def recv_msg(udp_socket):
    recv_data = udp_socket.recvfrom(1024)
    print('{info}:{msg}'.format(info=recv_data[1], msg=recv_data[0].decode('utf8')))


def main():
    udp_socket = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)
    udp_socket.bind(('127.0.0.1', 9090))

    while True:
        send_msg(udp_socket)
        recv_msg(udp_socket)


if __name__ == '__main__':
    main()