from concurrent.futures import ThreadPoolExecutor, as_completed, FIRST_COMPLETED, wait
import time
import threading


"""
    常规方式实例线程池
"""


class GetHtml(threading.Thread):
    def __init__(self, name, times):
        super().__init__(name=name)
        self._times = times

    def run(self):
        time.sleep(self._times)
        print("get page {} success".format(self._times))
        return self._times


def get_html(times):
    time.sleep(times)
    print("get page {} success".format(times))
    return times


if __name__ == '__main__':
    # max_workers表示线程池最大并发执行的工作线程数
    executor = ThreadPoolExecutor(max_workers=2)

    task1 = executor.submit(GetHtml("task1", 3).start())
    task2 = executor.submit(GetHtml("task2", 2).start())
    task3 = executor.submit(get_html, 3)
    task4 = executor.submit(get_html, 2)

    # done()用于判断线程任务是否执行完成
    print(task3.done())
    print(task4.done())
    
    # cancel()尝试中断等待队列中的某线程任务，若是被线程池执行则无法中断
    print(task2.cancel())
    
    time.sleep(5)
    
    # result()用于获得线程任务执行后的结果
    print(task3.result())
    print(task4.result())
    
    print(task3.done())
    print(task4.done())