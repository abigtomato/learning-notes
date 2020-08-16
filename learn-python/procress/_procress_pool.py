import multiprocessing
import time, os, random


"""
    进程池(池化思想就是不断复用已创建的进程，避免进程不断创建和销毁消耗大量资源)
"""


def worker(msg):
    start_time = time.time()
    print("开始执行：{msg}，进程号为：{pid}".format(msg=msg, pid=os.getpid()))
    time.sleep(random.random() * 2)
    end_time = time.time()
    print("结束执行：{msg}，执行时间：{time}".format(msg=msg, time=end_time-start_time))
    return msg


if __name__ == '__main__':
    # cpu_count()使进程池的容量根据cpu核数决定
    pool = multiprocessing.Pool(multiprocessing.cpu_count())
    # 
    for i in range(0, 10):
        # apply_async()开启新的进程
        result = pool.apply_async(worker, args=(i, ))
        # get()获得进程执行后的结果
        print("result {res}".format(res=result.get()))

    # close()禁止pool再接收进程
    pool.close()
    # join()等待池中的进程全部执行完毕，再执行主进程
    pool.join()