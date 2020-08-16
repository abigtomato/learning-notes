import multiprocessing
import time


"""
    进程池简化写法
"""


def get_html(num):
    time.sleep(num)
    print("success {num} page".format(num=num))
    return num


if __name__ == "__main__":
    pool = multiprocessing.Pool(multiprocessing.cpu_count())
    
    for result in pool.imap_unordered(get_html, [1, 5, 3]):
        print("{res} sleep success".format(res=result))