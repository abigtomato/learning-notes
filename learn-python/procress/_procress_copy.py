import multiprocessing
import time, os


"""
    多进程拷贝文件
"""


def copy_file(queue, file_name, old_folder_name, new_folder_name):
    print("进程{pid}将{file}从{old}拷贝到{new}".format(pid=os.getpid(), file=file_name, old=old_folder_name, new=new_folder_name))
    with open(old_folder_name + '/' + file_name, 'rb') as f:
        content = f.read()
    with open(new_folder_name + '/' + file_name, 'wb') as f:
        f.write(content)
    queue.put(file_name)


def main():
    old_folder_name = 'E:/usr/learn-python/02/thread'
    try:
        new_folder_name = old_folder_name + '_副本'
        os.mkdir(new_folder_name)
    except Exception:
        pass

    file_names = os.listdir(old_folder_name)
    print("要拷贝的文件列表：{list}".format(list=file_names))

    pool = multiprocessing.Pool(5)
    queue = multiprocessing.Manager().Queue()
    for file_name in file_names:
        pool.apply_async(copy_file, args=(queue, file_name, old_folder_name, new_folder_name))
    pool.close()
    
    file_names_len = len(file_names)
    copy_file_num = 0
    while True:
        queue.get()
        copy_file_num += 1
        print("\r拷贝进度：%.2f%%" % (copy_file_num*100/file_names_len), end='')
        if copy_file_num >= file_names_len:
            break


if __name__ == '__main__':
    main()