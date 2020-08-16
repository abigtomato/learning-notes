import asyncio
import time


"""
    asyncio实现并发，就需要多个协程来完成任务，每当有任务遇到阻塞操作的时候就await让其他协程继续工作
"""


now = lambda: time.time()


async def do_some_work(x):
    print("Waiting: {x}".format(x=x))
    await asyncio.sleep(x)
    return "Done after {x}s".format(x=x)


if __name__ == '__main__':
    start = now()

    # 创建协程对象
    coroutine1 = do_some_work(3)
    coroutine2 = do_some_work(2)
    coroutine3 = do_some_work(4)

    # 创建task的列表，然后将这些task注册到事件循环中
    tasks = [
        asyncio.ensure_future(coroutine1),
        asyncio.ensure_future(coroutine2),
        asyncio.ensure_future(coroutine3)
    ]
    loop = asyncio.get_event_loop()
    loop.run_until_complete(asyncio.wait(tasks))

    for task in tasks:
        print("Task ret: {ret}".format(ret=task.result()))
    
    print("Time: {time}".format(time=now() - start))