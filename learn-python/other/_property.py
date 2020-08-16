from datetime import date, datetime


class User(object):
    def __init__(self, name, birthday, info={}):
        self.name = name
        self.birthday = birthday
        self.__age = 0
        self.info = info

    @property
    def age(self):
        return datetime.now().year - self.birthday.year

    @age.setter
    def age(self, value):
        self.__age = value

    def __getattr__(self, item):
        return self.info[item]


if __name__ == '__main__':
    user = User('albert', date(year=1987, month=1, day=1), info={'school': 'mit'})

    print(user.age)

    user.age = 10
    print(user._User__age)

    print(user.school)