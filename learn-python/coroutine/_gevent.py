import gevent
import time
from gevent import monkey


"""
    使用gevent实现协程并发
"""


# 将程序中所有延时操作全部替换为gevent模块提供的延时操作
monkey.patch_all()


def func1(n):
    for i in range(n):
        # getcurrent()获取当前执行到此处的协程
        print(gevent.getcurrent(), i)
        
        # gevent.sleep(1)
        
        # gevent的协程触发到延时操则会切换到其他协程执行
        # 默认需要gevent模块提供的延时操作，使用monkey.patch_all()可自动替换
        time.sleep(1)


def func2(n):
    for i in range(n):
        print(gevent.getcurrent(), i)
        # gevent.sleep(1)
        time.sleep(1)


def func3(n):
    for i in range(n):
        print(gevent.getcurrent(), i)
        # gevent.sleep(1)
        time.sleep(1)


if __name__ == '__main__':
    # joinall()为列表中的所有协程都执行join操作
    # join()操作让当前线程等待指定协程执行完后再执行，此操作为延时操作，会触发协程执行
    gevent.joinall([
        # spawn()生成新协程并指定该协程的执行路径和要传入的参数
        gevent.spawn(func1, 5), 
        gevent.spawn(func2, 5), 
        gevent.spawn(func3, 5)
    ])