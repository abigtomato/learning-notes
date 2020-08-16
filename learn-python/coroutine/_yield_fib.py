import time


"""
    生成器：使用yield实现斐波那契数列的生成
"""


def create_fib(all_num):
    a, b = 0, 1
    current_num = 0
    while current_num < all_num:
        yield a
        a, b = b, a + b
        current_num += 1
    return 'end'


if __name__ == '__main__':
    fib = create_fib(10)
    while True:
        try:
            num = next(fib)
            print(num)
        except StopIteration as ret:
            print(ret.value)
            break