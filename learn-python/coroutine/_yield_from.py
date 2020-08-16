from itertools import chain
import time


"""
    yield from语法
"""


def my_chain(*args, **kwargs):
    for my_iterable in args:
        # for value in my_iterable:
        #     yield value
        yield from my_iterable


def g1(iterable):
    yield range(10)


def g2(iterable):
    yield from range(10)


if __name__ == '__main__':
    for value in my_chain(["hadoop", "spark", "flink"],
            {"albet": "http://www.baidu.com", 
             "lily": "http://www.google.com"}, range(10)):
        print("my_chain: {val}".format(val=value))
    
    for value in g1(range(10)):
        print("g1: {val}".format(val=value))
    
    for value in g2(range(10)):
        print("g2: {val}".format(val=value))