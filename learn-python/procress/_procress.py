import multiprocessing
import time


"""
    多进程编程(进程是资源分配的单位，是代码+所需资源的总和)
"""


def get_html(num):
    time.sleep(num)
    print("success {num} page".format(num=num))
    return num


def get_url(num):
    time.sleep(num)
    print("success {num} url".format(num=num))
    return num


if __name__ == '__main__':
    """
        进程状态：
            1.新建---启动--->就绪<---调度--->运行---结束--->死亡
            2.运行---等待条件--->等待(阻塞)---满足条件--->就绪
            就绪态：运行条件满足，等待cpu执行
            执行态：cpu正在执行
            等待态：等待某些条件满足(例如sleep()，等待睡眠时间结束)
    """
    progress_01 = multiprocessing.Process(target=get_url, args=(2, ))
    progress_02 = multiprocessing.Process(target=get_html, args=(2, ))

    """
        进程是逻辑代码+物理资源的总和，当一个进程创建出新的子进程时：
        1.若子进程会对内存资源中的数据进行修改(执行写操作)，那么根据操作系统写时复制的特性，子进程会拷贝父进程的代码
          和内存(开辟属于自己的内存空间，拷贝父进程内存数据)；
        2.若子进程不会对内存资源进行修改，那么子进程会复用父进程的代码和资源，在父进程的空间中执行自己的逻辑。
    """
    progress_01.start()
    progress_02.start()
    print("进程id: {pid}".format(pid=progress_01.pid))
    print("进程id: {pid}".format(pid=progress_02.pid))
    
    # 等待指定进程执行结束再执行主进程
    progress_01.join()
    progress_02.join()
    print("main progress")