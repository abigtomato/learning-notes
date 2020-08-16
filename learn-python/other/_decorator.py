import time


# 闭包实现装饰器
def decorator(func):
    def wrapper(*args, **kw):
        print(time.time())
        func(*args, **kw)
    return wrapper


# AOP编程，动态的添加代理以增强方法
@decorator
def func1(func_name):
    print("decorator......" + func_name)


@decorator
def func2(*args):
    res = [arg for arg in args]
    print(res)


@decorator
def func3(*args, **kw):
    print([arg for arg in args])
    print(kw)


if __name__ == '__main__':
    func1("hadoop")
    func2("spark core", "spark sql", "spark streaming", "spark mllib")
    func3("kafka", "scala", a=1, b=2, c=3)
