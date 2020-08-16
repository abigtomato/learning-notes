import gevent
from gevent import monkey
import urllib


"""
    使用gevent协程并发下载图片
"""


monkey.patch_all()


def downloader(img_name, url):
    req = urllib.request.urlopen(url)
    # 网络请求时会阻塞等待，这时等待时间就可以被拿出来利用
    # 此时会在当前线程汇总衍生出一条新的执行路径，也就是一条协程，会利用阻塞等待时间去执行别的逻辑
    img_content = req.read()

    with open('./' + img_name, 'wb') as f:
        f.write(img_content)


def main():
    # 协程就是利用了代码阻塞等待的时间去执行其他任务
    gevent.joinall([
        gevent.spawn(downloader, '1.jpg', ''),
        gevent.spawn(downloader, '2.jpg', ''),
        gevent.spawn(downloader, '3.jpg', '')
    ])


if __name__ == '__main__':
    main()