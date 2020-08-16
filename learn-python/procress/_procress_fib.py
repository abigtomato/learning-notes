from concurrent.futures import ThreadPoolExecutor, ProcessPoolExecutor, as_completed
import time


"""
    多线程多进程执行计算密集型任务比较
"""


# 斐波那契函数
def fib(num):
    if num <= 2:
        return 1
    return fib(num-1) + fib(num-2)


if __name__ == '__main__':
    """1.使用多线程计算斐波那契数"""
    with ThreadPoolExecutor(3) as executor:
        all_task = [executor.submit(fib, (num)) for num in range(25, 40)]

        start_time = time.time()
        for future in as_completed(all_task):
            data = future.result()
            print("thread result: {data}".format(data=data))
        end_time = time.time()
        print("thread time is: {time}".format(time=end_time-start_time))

    """2.使用多进程计算斐波那契数"""
    with ProcessPoolExecutor(3) as executor:
        all_task = [executor.submit(fib, (num)) for num in range(25, 40)]

        start_time = time.time()
        for future in as_completed(all_task):
            data = future.result()
            print("process result: {data}".format(data=data))
        end_time = time.time()
        print("process time is: {time}".format(time=time.time() - start_time))