import time


"""
    yield from统计示例
"""


final_result = {}


# 子生成器
def sales_sum(key):
    total = 0
    nums = []
    while True:
        # 接收调度器传递的值
        x = yield
        print("{key}销量: {num}".format(key=key, num=x))
        # 子生成器终止条件
        if not x:
            break
        total += x
        nums.append(x)
    # 返回结果给委托生成器
    return total, nums


# 委托生成器
def middle(key):
    while True:
        # 使用yield from建立调度器和子生成器的双向通道
        # 同时接收子生成器的返回值
        final_result[key] = yield from sales_sum(key)
        print("{key}销量统计完成".format(key=key))


# 调度器
def main():
    data_sets = {
        "Hadoop": [1200, 1500, 3000],
        "Spark": [28, 55, 98, 108],
        "Flink": [280, 560, 778, 70]
    }

    for key, data_set in data_sets.items():
        print("开始统计: {key}".format(key=key))
        # 委托生成器对象
        m = middle(key)
        # 预激活子生成器(直接传递给子生成器)
        m.send(None)
        for val in data_set:
            # 向子生成器传值
            m.send(val)
        # 传None值结束子生成器
        m.send(None)
    print("统计结果: {ret}".format(ret=final_result))


if __name__ == '__main__':
    main()