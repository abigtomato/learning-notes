import time 
import greenlet


"""
    使用greenlet实现协程并发
"""


def task_01():
    while True:
        time.sleep(1)
        print("---1---")
        # 切换到task_02中执行
        gr_2.switch()        


def task_02():
    while True:
        time.sleep(1)
        print("---2---")
        # 切换到task_01中执行
        gr_1.switch()


gr_1 = greenlet.greenlet(run=task_01)
gr_2 = greenlet.greenlet(run=task_02)

# 到task_01中执行 
gr_1.switch()
