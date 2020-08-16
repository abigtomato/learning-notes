import threading
import time


class HtmlSpider(threading.Thread):
    def __init__(self, url, sem):
        super().__init__()
        self.url = url
        self._sem = sem

    def run(self):
        time.sleep(2)
        print(self.url)
        self._sem.release()


class UrlProducer(threading.Thread):
    def __init__(self, sem):
        super().__init__()
        self._sem = sem

    def run(self):
        for i in range(20):
            self._sem.acquire()
            HtmlSpider('http://baidu.com/{num}'.format(num=i), self._sem).start()


if __name__ == '__main__':
    sem = threading.Semaphore(3)
    UrlProducer(sem).start()