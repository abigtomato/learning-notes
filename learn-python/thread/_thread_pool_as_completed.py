from concurrent.futures import ThreadPoolExecutor, as_completed, FIRST_COMPLETED, wait
import time
import threading


"""
    executor.map的方式执行task
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
    executor = ThreadPoolExecutor(max_workers=2)

    urls = [3, 2, 4]
    for data in executor.map(get_html, urls):
        print("result: {data}".format(data=data))