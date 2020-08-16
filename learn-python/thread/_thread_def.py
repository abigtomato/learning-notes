import threading
import time

"""
    传递执行函数的方式实现多线程(线程是操作系统调度的单位)
"""

total = 100


def get_html(html=None):
    # 多线程共享全局变量
    global total
    total += 1

    print("get_html start")
    time.sleep(2)
    print("get_html end")


def get_url(url=None):
    print("get_url start")
    time.sleep(2)
    print("get_url end")

    # 多线程共享全局变量
    print("total:{total}".format(total=total))


if __name__ == '__main__':
    # 传递执行函数的方式实现多线程
    thread1 = threading.Thread(target=get_html, args=("html", ))
    thread2 = threading.Thread(target=get_url, args=("url", ))

    # setDaemon(True)将当前线程设置为守护线程(等待主线程执行完毕后立即kill)
    thread1.setDaemon(True)
    thread2.setDaemon(True)

    start = time.time()
    thread1.start()
    thread2.start()
    print("当前线程信息:{}".format(threading.enumerate()))
    # 设置主线程阻塞等待子线程执行完毕(相当于将主线程的执行路径join连接到了子线程之后)
    thread1.join()
    thread2.join()
    end = time.time()

    print("{time}".format(time=end - start))