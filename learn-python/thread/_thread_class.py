import threading
import time


"""
    继承Thread类的方式实现多线程
"""


class GetHtml(threading.Thread):
    def __init__(self, name):
        super().__init__(name=name)

    def run(self):
        print("GetHtml start")
        time.sleep(2)
        print("GetHtml end")


class GetUrl(threading.Thread):
    def __init__(self, name):
        super().__init__(name=name)

    def run(self):
        print("GetUrl start")
        time.sleep(2)
        print("GetUrl end")


if __name__ == "__main__":
    # 继承Thread类的方式实现多线程
    thread1 = GetHtml("get_html")
    thread2 = GetUrl("get_url")

    start = time.time()
    thread1.start()
    thread2.start()
    thread1.join()
    thread2.join()
    end = time.time()

    print("{time}".format(time=end - start))