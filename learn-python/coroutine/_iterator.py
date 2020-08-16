# from collections import Iterable, Iterator


"""
    迭代器(特点就是占用极小的内存空间生成需要的数据)
"""


class Classmate(object):
    def __init__(self):
        self._names = list()
        self._current_num = 0
    
    def add(self, name):
        self._names.append(name)
    
    # 存在__iter__方法的对象是可迭代对象
    # 将当前对象做为循环对象操作时会指向此方法返回迭代器对象
    def __iter__(self):
        # 返回的对象就是迭代器
        return ClassIter(self)


class ClassIter(object):
    def __init__(self, obj):
        self._obj = obj
        self._current_num = 0

    # 存在__next__方法的对象是迭代器对象
    # 在使用for语法操作迭代器时每次循环都会自动调用此方法
    def __next__(self):
        if self._current_num < len(self._obj._names):
            elem = self._obj._names[self._current_num]
            self._current_num += 1
            return elem
        else:
            raise StopIteration


if __name__ == '__main__':
    classmate = Classmate()
    classmate.add('spark sql')
    classmate.add('spark streaming')
    classmate.add('spark mllib')

    # print("判断classmate是否是可迭代对象: {}".format(isinstance(classmate, Iterable)))
    # print("判断classmate是否是迭代器: {}".format(isinstance(classmate, Iterator)))

    for elem in classmate:
        print(elem)