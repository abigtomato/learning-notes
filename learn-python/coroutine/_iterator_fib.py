"""
    迭代器应用：生成斐波那契数列
"""


class Fibonacci(object):
    def __init__(self, all_num):
        self._all_num = all_num
        self._current_num = 0
        self._a = 0
        self._b = 1
    
    def __iter__(self):
        return self
    
    def __next__(self):
        if self._current_num < self._all_num:
            res = self._a

            self._a, self._b = self._b, self._a + self._b
            self._current_num += 1

            return res
        else:
            raise StopIteration


if __name__ == '__main__':
    fibo = Fibonacci(10)
    print(list(fibo))