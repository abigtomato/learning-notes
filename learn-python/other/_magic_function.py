import time


"""
    魔法函数
"""


class Company(object):
    def __init__(self, employee_lit):
        self.employee = employee_lit

    # 魔法函数，当对象进行迭代切片等操作时，若不具备这些特性，解释器则调用该函数中的逻辑
    def __getitem__(self, item):
        return self.employee[item]

    # 让对象具备长度的特性
    def __len__(self):
        return len(self.employee)

    # 让对象可被字符串描述
    def __str__(self):
        return ','.join(self.employee)

    # 当对象被垃圾回收时执行
    def __del__(self):
        print("delete .....")
    

class MyNumber(object):
    def __init__(self, num):
        self.__num = num

    # 让对象具备数学计算的特性
    def __abs__(self):
        return abs(self.__num)


class MyVector(object):
    def __init__(self, x, y):
        self.__x = x
        self.__y = y

    # 让对象具备数学计算的特性
    def __add__(self, other):
        return MyVector(self.__x + other.__x, self.__y + other.__y)

    def __str__(self):
        return '({x}, {y})'.format(x=self.__x, y=self.__y)


if __name__ == '__main__':
    company = Company(['spark', 'kafka', 'scala'])

    # 测试__getitem__
    for elem in company:
        print(elem)
    print(company[: 2])

    # 测试__len__
    print(len(company))

    # 测试__str__
    print(company)

    # 测试__abs__
    print(abs(MyNumber(-1)))

    # 测试__add__
    first_vec = MyVector(1, 2)
    second_vec = MyVector(2, 3)
    print(first_vec + second_vec)

    # company对象引用计数减1，当为0时触发垃圾回收
    del company