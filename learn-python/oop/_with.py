# 异常处理
def exe_try():
    try:
        print("exception ......")
        raise KeyError
    except KeyError as e:
        print("key error" + e)
        return 1
    else:
        print("index error")
        return 2
    finally:
        print("finally")
        return 3


class Sample:
    # 前置执行
    def __enter__(self):
        print("获取资源 。。。。。。")
        return self

    # 后置执行
    def __exit__(self, exc_type, exc_val, exc_tb):
        print("释放资源 。。。。。。")

    # 业务逻辑
    def do_something(self):
        print("代码逻辑 。。。。。。")


# with上下文管理器协议
with Sample() as sample:
    sample.do_something()


import contextlib


# contextlib和生成器的方式简化上下文管理器
@contextlib.contextmanager
def file_open(file_name):
    print("file open")
    yield {}
    print("file end")


with file_open("bobby.txt") as f_opened:
    print("file processing")