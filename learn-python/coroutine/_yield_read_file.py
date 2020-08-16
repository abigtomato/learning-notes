import time


"""
    生成器：读取一行大文件示例
"""


def my_read_lines(f, newline):
    buff = ''
    while True:
        while newline in buff:
            pos = buff.index(newline)
            yield buff[: pos]
            buff = buff[pos + len(newline): ]
        chunk = f.read(4096)
    
        if not chunk:
            yield buff
            break
        buff += chunk

    
if __name__ == '__main__':
    with open('E:/usr/learn-python/02/coroutine/data/input.txt', 'r') as f:
        for line in my_read_lines(f, '{|}'):
            print(line)