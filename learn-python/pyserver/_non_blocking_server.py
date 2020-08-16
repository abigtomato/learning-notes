import time
import socket


"""
    非阻塞实现单进程，单线程并发服务器
"""


def main():
    tcp_server_socket = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    tcp_server_socket.setsockopt(socket.SOL_SOCKET, socket.SO_REUSEADDR, 1)
    tcp_server_socket.bind(("127.0.0.1", 7890))
    tcp_server_socket.listen(128)

    # 设置为非阻塞
    tcp_server_socket.setblocking(False)
    client_socket_list = []
    
    while True:
        time.sleep(1)

        try:
            # 此处不会阻塞等待连接建立，无连接建立会抛出异常
            client_socket, client_addr = tcp_server_socket.accept()
        except Exception as ret:
            print("未建立连接: {ret}".format(ret=ret))
        else:
            print("建立连接: {info}".format(info=client_addr))
            client_socket.setblocking(False)
            client_socket_list.append(client_socket)

        for client_socket in client_socket_list:
            try:
                # 此处不会阻塞等待数据接收，没有接收到数据会抛出异常
                recv_data = client_socket.recv(1024)
            except Exception as ret:
                print("未接收数据: {ret}".format(ret=ret))
            else:
                if recv_data:
                    print("接收的数据: {data}".format(data=recv_data))
                else:
                    client_socket_list.remove(client_socket)
                    client_socket.close()
               

if __name__ == '__main__':
    main()