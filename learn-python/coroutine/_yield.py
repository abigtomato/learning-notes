import time

"""
    生成器(特点是让函数在任意点暂停或恢复执行)
"""

def demo():
    print("第一次yield之前的代码逻辑")
    # yield后面跟的语句是产出值，若yield做为表达式右边，那么左边则是生成器接收的值
    html = yield 'http://www.baidu.com'
    print("第一次从send接收到的值", html)

    print("\n第二次yield之前的代码逻辑")
    num = yield 2
    print("第二次从send接收到的值", num)

    print("\n第三次yield之前的代码逻辑")
    yield 3

    print("\n第四次yield之前的代码逻辑")
    try:
        yield 4
    except Exception as e:
        print("此处捕获到异常", e)

    return 'bobby'


if __name__ == '__main__':
    # 生成器对象
    demo = demo()

    # next使生成器执行到第一个yield时暂停整个函数
    url = next(demo)
    # url为第一个yield返回的值，也就是yield关键字右边的语句
    print("第一次yield的返回值：", url)

    # send往生成器中传递值(第一次调用send则由生成器的第一个yield接收)
    # 并使生成器继续执行到第二个yield，返回值后暂停生成器
    print("第二次yield的返回值：", demo.send("<html></html>"))

    # 第二次向生成器中send值会由第二个yield接收，并继续执行到下一个yield暂停
    print("第三次yield的返回值：", demo.send("<body></body>"))

    # 再次调用next会从上一次暂停的yield之后执行代码，直到下一个yield关键字，之后再有类似的逻辑就以此类推
    print("第四次yield的返回值：", next(demo))

    # 向当前暂停的yield抛出异常
    try:
        demo.throw(Exception, "download error")
    except StopIteration:
        pass

    # 关闭生成器
    demo.close()