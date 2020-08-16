from concurrent.futures import ThreadPoolExecutor, as_completed, FIRST_COMPLETED, wait
import time
import threading


"""
    as_completed()方式执行线程池的任务
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


if __name__ == "__main__":
    executor = ThreadPoolExecutor(max_workers=2)

    urls = [3, 2, 4]
    all_task = [executor.submit(get_html, url) for url in urls]
    # wait()使主线程阻塞，等待线程池中task结束再执行
    # return_when参数指定等待的形式，默认等待全部执行完，FIRST_COMPLETED表示等待第一个task结束就中断阻塞
    wait(all_task, return_when=FIRST_COMPLETED)

    # future表示未来结果对象
    for future in as_completed(all_task):
        # 通过未来对象获取task执行后的结果
        data = future.result()
        print("result: {data}".format(data=data))