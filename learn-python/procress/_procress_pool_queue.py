import time
import multiprocessing


"""
    进程池中进程的通信
"""


def download_from_web(queue):
    data = ['hadoop', 'hive', 'hbase', 'spark']
    for temp in data:
        # full()判断队列容量是否满了
        if not queue.full():
            # put()入队操作，若队列容量满了则阻塞
            queue.put(temp)
    time.sleep(2)


def analysts_data(queue):
    time.sleep(2)
    waitting_analysis_data = list()
    while True:
        # get()出队操作，若队列为空则阻塞
        data = queue.get()
        waitting_analysis_data.append(data)
        # empty()判断队列是否为空
        if queue.empty():
            break
    print(waitting_analysis_data)


if __name__ == '__main__':
    # 进程池中进程的通信，需要使用Manager模块下提供的消息队列"""
    queue = multiprocessing.Manager().Queue(10)
    pool = multiprocessing.Pool(2)
    
    pool.apply_async(download_from_web, args=(queue, ))
    pool.apply_async(analysts_data, args=(queue, ))
    
    pool.close()
    pool.join()