import threading


"""
    等待唤醒机制
"""


class XiaoAi(threading.Thread):
    def __init__(self, name, cond):
        super().__init__(name=name)
        self.cond = cond

    def run(self):
        # 使用with上下文管理器自动加锁释放锁
        with self.cond:
            # wait()使持有当前cond条件的线程陷入等待
            self.cond.wait()
            print("{} : 在 ".format(self.name))
            # notify()唤醒当前持有cond条件的线程
            self.cond.notify()

            self.cond.wait()
            print("{} : 好啊 ".format(self.name))
            self.cond.notify()

            self.cond.wait()
            print("{} : 君住长江尾 ".format(self.name))
            self.cond.notify()

            self.cond.wait()
            print("{} : 共饮长江水 ".format(self.name))
            self.cond.notify()

            self.cond.wait()
            print("{} : 此恨何时已 ".format(self.name))
            self.cond.notify()

            self.cond.wait()
            print("{} : 定不负相思意 ".format(self.name))
            self.cond.notify()


class TianMao(threading.Thread):
    def __init__(self, name, cond):
        super().__init__(name=name)
        self.cond = cond

    def run(self):
        # 为持有此cond条件的线程加锁
        self.cond.acquire()

        print("{} : 小爱同学 ".format(self.name))
        self.cond.notify()
        self.cond.wait()

        print("{} : 我们来对古诗吧 ".format(self.name))
        self.cond.notify()
        self.cond.wait()

        print("{} : 我住长江头 ".format(self.name))
        self.cond.notify()
        self.cond.wait()

        print("{} : 日日思君不见君 ".format(self.name))
        self.cond.notify()
        self.cond.wait()

        print("{} : 此水几时休 ".format(self.name))
        self.cond.notify()
        self.cond.wait()

        print("{} : 只愿君心似我心 ".format(self.name))
        self.cond.notify()
        self.cond.wait()

        # 释放持有此cond条件的线程
        self.cond.release()


if __name__ == '__main__':
    cond = threading.Condition()
    
    xiaoai = XiaoAi('xiaoai', cond)
    tianmao = TianMao('tianmao', cond)
    xiaoai.start()
    tianmao.start()