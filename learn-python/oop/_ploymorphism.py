class Cat(object):
    def say(self):
        print("cat ......")


class Dog(object):
    def say(self):
        print("dog ......")


class Duck(object):
    def say(self):
        print("duck ......")


class MyList(object):
    def __init__(self, employee):
        self.employee = employee

    def __getitem__(self, item):
        return self.employee[item]


if __name__ == '__main__':
    # 多态，animal引用可随意指向任意实现了say()方法的类
    animal_list = [Cat, Dog, Duck]
    for animal in animal_list:
        animal().say()

    # 列表连接列表，只要是可迭代的对象就可以通过extend连接
    list_1 = [i**2 for i in range(10)]
    list_2 = [i**3 for i in range(10)]
    list_1.extend(list_2)
    print(list_1)

    # 列表连接集合
    set_1 = set()
    set_1.add(1)
    set_1.add(2)
    set_1.add(3)
    list_1.extend(set_1)
    print(list_1)

    # 列表连接实现了可迭代特性的类
    my_list = MyList([5, 4, 3, 2, 1])
    list_1.extend(my_list)
    print(list_1)
