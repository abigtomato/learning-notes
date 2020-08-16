import multiprocessing


"""
    进程间通信的容器
"""


def add_dict(p_dict, key, value):
    p_dict[key] = value


if __name__ == '__main__':
    # 通信字典
    p_dict = multiprocessing.Manager().dict()
    process = multiprocessing.Process(target=add_dict, args=(p_dict, 'big data', 'spark'))
    
    process.start()
    process.join()
   
    print(p_dict)