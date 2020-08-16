import threading


"""
    重入锁
"""


class Reenter(threading.Thread):
    def __init__(self, thread_name, rlock):
        super().__init__(name=thread_name)
        self.rlock = rlock

    def run(self):
        for i in range(10000):
            self.rlock.acquire()
            print(i)
            self.rlock.release()


if __name__ == '__main__':
    # 重入锁
    rlock = threading.RLock()

    reenter = Reenter("reenter", rlock)
    reenter.start()
    reenter.join()