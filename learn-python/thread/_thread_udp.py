import socket
import threading


"""
    多线程UDP通信
"""


def recv_msg(udp_socket):
    while True:
        recv_data = udp_socket.recvfrom(1024)
        print(recv_data.decode('utf8'))


def send_msg(udp_socket, dest_ip, dest_port):
    while True:
        send_data = input("")
        udp_socket.sendto(send_data.encode('utf8'), (dest_ip, dest_port))


def main():
    udp_socket = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)
    udp_socket.bind(("127.0.0.1", 7890))

    t_send = threading.Thread(target=send_msg, args=(udp_socket, "127.0.0.1", 7890))
    t_recv = threading.Thread(target=recv_msg, args=(udp_socket))
    t_send.start()
    t_recv.start()


if __name__ == "__main__":
    main()