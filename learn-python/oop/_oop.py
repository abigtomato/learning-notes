# Python中一切皆为对象
class Human(object):
    __height = 10

    def __init__(self, name, age):
        self.__name = name
        self.__age = age

    def get_name(self):
        return self.__name


class Chinese(object):
    def __init__(self):
        super().__init__()


# 继承自Human
class Student(Human, Chinese):
    # 类变量
    __sum = 0

    # 构造器
    def __init__(self, name, age, school='mit', cls='20180304'):
        # 调用父类构造器，存在多继承和多层继承情况下，根据mro算法的顺序调用父类构造器
        super(Student, self).__init__(name, age)

        # "self.xxx"表示实例变量，"__"前缀表示私有化
        self.__school = school
        self.__cls = cls

        # __class__为内置属性，表示类的引用
        self.__class__.__sum += 1

    # 实例方法
    def stu_print(self):
        # self表示当前对象的引用
        # __dict__以字典格式查看当前对象所有属性
        print(self.__dict__)
        print(self.__class__.__sum)

        # 调用父类实例方法
        name = super(Student, self).get_name()
        print(name)

    # 类方法
    @classmethod
    def plus_sum(cls):
        # cls表示类的引用
        cls.__sum += 1
        print(cls.__sum)

    # 静态方法
    @staticmethod
    def static_add(x, y):
        print(x, y)


class Date(object):
    def __init__(self, year, month, day):
        self.year = year
        self.month = month
        self.day = day

    def tomorrow(self):
        self.day += 1

    # 涉及到类操作时使用类方法，如此处的字符串转Date对象
    @classmethod
    def from_str(cls, date_str):
        year, month, day = tuple(date_str.split('-'))
        return cls(int(year), int(month), int(day))

    # 涉及到与对象和类无关操作时使用静态方法，如此处的日期格式校验
    @staticmethod
    def valid_str(date_str):
        year, month, day = tuple(date_str.split('-'))
        if int(year) > 0:
            return True
        else:
            return False

    def __str__(self):
        return '{year}/{month}/{day}'.format(year=self.year, month=self.month, day=self.day)


if __name__ == '__main__':
    stu = Student('albert', 18)
    stu.stu_print()

    Student.plus_sum()
    Student.static_add(1, 2)

    # 动态语言特性，相当于为当前对象新添加了一个属性"__name"
    stu.__name = 'lily'
    print(stu.__name)

    # 私有属性会添加"_类名"的前缀来进行隐藏
    print(stu._Student__school)

    print(stu.get_name())

    # 类和实例属性查找顺序
    print(Student.__mro__)

    # 对象的自省机制
    print(stu.__dict__)
    print(Student.__dict__)

    # 对象动态扩充属性
    stu.__dict__['school_name'] = 'MIT'
    print(stu.school_name)
