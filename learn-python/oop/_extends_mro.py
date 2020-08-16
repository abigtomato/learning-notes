import time


"""
    多继承中的MRO顺序
"""


class Parent(object):
    def __init__(self, name, *args, **kwargs):
        print("Parent的__init__被调用")
        self.__name = name

    
class Son1(Parent):
    def __init__(self, name, age, *args, **kwargs):
        print("Son1的__init__被调用")
        self.__age = age
        super().__init__(name, *args, **kwargs)
  

class Son2(Parent):
    def __init__(self, name, gender, *args, **kwargs):
        print("Son2的__init__被调用")
        self.__gender = gender
        super().__init__(name, *args, **kwargs)


class Grandson(Son1, Son2):
    def __init__(self, name, age, gender):
        print("Grandson的__init__被调用")
        super().__init__(name, age, gender)


if __name__ == '__main__':
    # 根据c3算法计算出的mro调用顺序(多继承下的super().__init__()方法调用顺序)
    grandson = Grandson("albert", 18, "hello")
    print(Grandson.__mro__)