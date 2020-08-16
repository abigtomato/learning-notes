def add(a, b):
    # += 会将拼接结果赋值给a
    a += b
    return a


class Error(object):
    # 默认值预先创建好，若是没有参数传入，则当前对象存在指向默认值的引用
    def __init__(self, name, elems=[]):
        self.name = name
        self.elems = elems

    def add(self, item):
        self.elems.append(item)

    def rem(self, item):
        self.elems.remove(item)


if __name__ == '__main__':
    a = [1, 2, 3]
    b = [4, 5, 6]
    c = add(a, b)
    print(c)
    # 引用传递，改变了a的值
    print("{a}:{b}".format(a=a, b=b))

    # err1对象没有传递参数2，则存在一个引用指向默认值
    err1 = Error('err1')
    err1.add('hello1')
    print(err1.elems)

    # err2对象也没有传递参数2，也存在一个引用指向err1指向的默认值，所有两个对象的elems属性相同(指向同一块内存)
    err2 = Error('err2')
    print(err2.elems)