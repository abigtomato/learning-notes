import re


"""
    python的re模块使用
"""


def add(temp):
    strNum = temp.group()
    num = int(strNum) + 1
    return str(num)


def main():
    # 匹配
    ret = re.match(r"^[a-zA-Z0-9_]{4,20}@(163|126)\.com$", "18895672556@163.com")
    print("email result: {ret}".format(ret=ret.group()))
    
    ret = re.match(r"^(http|https)://([a-zA-Z0-9_]{3,18}\.?)+\.(cn|com)$", "http://www.baidu.com")
    print("url result: {ret}".format(ret=ret.group()))

    # 查找全部
    ret = re.findall(r"\d+", "阅读数: 9999，点赞数: 1000")
    print("num result: {ret}".format(ret=ret))

    # 替换
    ret = re.sub(r"\d+", add, "python = 997")
    print("sub result: {ret}".format(ret=ret))


if __name__ == '__main__':
    main()