import time


"""
    生成器：使用yield实现协程并发
"""


def task_01():
    while True:
        time.sleep(1)
        print("---1---")
        yield


def task_02():
    while True:
        time.sleep(1)
        print("---2---")
        yield

    
if __name__ == '__main__':
    t_1 = task_01()
    t_2 = task_02()

    while True:
        next(t_1)
        next(t_2)