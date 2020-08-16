import time
import asyncio
import functools


"""
    event_loop事件循环: 
        程序开启一个无限的循环，程序员会把一些函数注册到事件循环上，当满足事件触发条件时，自动调用相应的协程函数。
    coroutine协程: 
        协程对象，指一个使用async关键字声明的函数，调用它不会立即执行，而是返回一个协程对象，协程对象需要注册到事件循环，由事件循环调用。
    task任务: 
        一个协程对象就是一个可以被挂起的函数，task则是对协程进一步封装，其中包含任务的各种状态。
    future未来结果: 
        代表将来执行或没有执行的任务的结果，它和task上没有本质的区别
    async/await关键字: 
        python3.5用于定义协程的关键字，async定义一个协程，await用于挂起阻塞的函数。
"""


now = lambda : time.time()


# 通过async关键字定义一个协程(coroutine)，协程也是一种对象
# 协程不能直接运行，需要加入到事件循环(loop)中，由后者在适当的时候调用协程
async def do_some_work(x):
    print("Waiting: {x}".format(x=x))
    # 使用await可以针对耗时的操作将当前协程挂起，就像生成器里的yield一样，函数暂停让出控制权，下次得到控制权后从暂停处开始执行
    # 协程遇到await，事件循环将会挂起该协程，执行别的协程，直到其他的协程也挂起或者执行完毕，再从断点开始执行
    # 耗时的操作一般是一些IO操作，例如网络请求，文件读取等，此处使用asyncio.sleep()函数来模拟IO操作，协程的目的就是让这些IO操作异步化
    await asyncio.sleep(x)
    return "Done after {x}s".format(x=x)


# 协程do_some_work结束会调用回调函数，之后通过参数future获取协程执行的结果
def callback(url, future):
    print("{url} Callback: {ret}".format(url=url, ret=future.result()))


if __name__ == '__main__':
    start = now()
    coroutine = do_some_work(2)
    
    # asyncio.get_event_loop()创建一个事件循环
    loop = asyncio.get_event_loop()

    # # 协程对象不能直接运行，在注册事件循环的时候，其实是run_until_complete方法将协程包装成为了一个任务(task)对象
    # # 所谓task对象是Future类的子类，保存了协程运行后的状态，用于未来获取协程的结果，loop.create_task()应用将协程对象创建成task任务
    # task = loop.create_task(coroutine)
    
    # 将协程对象包装为task任务
    task = asyncio.ensure_future(coroutine)
    # add_done_callback()为task绑定回调函数，在task执行完毕的时候可以获取执行的结果，回调的最后一个参数是future对象，通过该对象可以获取协程返回值
    # 如果回调需要多个参数，可以通过functools.partial()偏函数导入
    task.add_done_callback(functools.partial(callback, "http://www.baidu.com"))

    # 使用run_until_complete()将协程注册到事件循环中，并启动事件循环
    loop.run_until_complete(task)

    print("result: {ret}".format(ret=task.result()))
    print("TIME: {time}".format(time=now() - start))