"""
    函数InsFunc是函数ExFunc的内嵌函数，并且是ExFunc函数的返回值
    内嵌函数InsFunc中引用到外层函数中的局部变量sum
    分别由不同的参数调用ExFunc函数得到的函数时(myFunc()，myAnotherFunc())，得到的结果是隔离的，也就是说每次调用ExFunc函数后都将生成并保存一个新的局部变量sum
    这里ExFunc函数返回的就是闭包
"""
def ExFunc(n):
    sum = n
    def InsFunc():
        return sum + 1
    return InsFunc


f1 = ExFunc(10)
f2 = ExFunc(20)
print(f1(), f2())


"""
    如果在一个内部函数里(adder(y)就是这个内部函数)
    对在外部作用域(但不是在全局作用域)的变量进行引用(x就是被引用的变量，x在外部作用域addx里面，但不在全局作用域里)
    则这个内部函数adder就是一个闭包
    闭包 = 函数块 + 函数依赖的环境(adder就是函数块，x就是环境)
"""
def addx(x):
    def adder(y):
        return x + y
    return adder


print(addx(3)(4))
print(type(addx(3)))
print(addx(3).__name__)


"""
    闭包中是不能修改外部作用域的局部变量的
    这里foo1()中的m是闭包的局部变量
"""
def foo():
    m = 0
    def foo1():
        m = 1
        print(m)

    print(m)
    foo1()
    print(m)


foo()


"""
    1.在执行代码c = fun1()时，解释器会导入全部的闭包函数体fun2()来分析其的局部变量
    2.解释器规则指定所有在赋值语句左面的变量都是局部变量，而在闭包fun2()中，变量a在赋值符号"="的左面，被python认为是fun2()中的局部变量
    3.再接下来执行调用后程序运行至a = a + 1时，因为先前已经把a归为fun2()中的局部变量，所以python会在fun2()中去找在赋值语句右面的a的值，结果找不到，就会报错
    解决方法:
        使用外部局部变量时添加nonlocal，显示声明a不是闭包的局部变量
"""
def fun1():
    a = 1
    def fun2():
        nonlocal a
        a = a + 1
        return a
    return fun2


print(fun1()())


"""
    用途:
        当闭包执行完后，仍然能够保持住当前的运行环境
        如果希望函数的每次执行结果，都是基于这个函数上次的运行结果
    1.假设棋盘大小为50*50，左上角为坐标系原点(0,0)，需要一个函数，接收2个参数，分别为方向(direction)，步长(step)，该函数控制棋子的运动；
    2.棋子运动的新的坐标除了依赖于方向和步长以外，当然还要根据原来所处的坐标点，用闭包就可以保持住这个棋子原来所处的坐标；
    3.这里的pos(也就是闭包的返回值)就是每次执行闭包后保留的结果。
"""
origin = (0, 0)     # 原始坐标
legal_x = (0, 50)   # x轴范围
legal_y = (0, 50)   # y轴范围


def create(pos = origin):
    # direction方向，step步长
    def player(direction, step):
        nonlocal pos
        pos = (pos[0] + direction[0] * step, pos[1] + direction[1] * step)
        # 闭包返回值，下次闭包调用会基于本次pos的值再做计算
        return pos
    return player


player = create()
print(player((1, 0), 10))
print(player((0, 1), 20))
print(player((-1, 0), 10))
print(player((0, -1), 20))
