from functools import reduce

# map对每个元素进行操作
list_x = [1, 2, 3, 4, 5, 6]
list_y = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]
res = map(lambda x, y: x * x + y, list_x, list_y)
print(list(res))

# reduce对所有元素进行聚合
res = reduce(lambda x, y: x + y, list_x)
print(res)

# filter对元素进行过滤
list_z = [1, 0, 1, 2, 0, 1, 3, 4, 1, 1, 5]
res = filter(lambda x: True if x is 1 else False, list_z)
print(list(res))


def outer_func(item):
    return item ** 2


# 列表生成式
res = [outer_func(i) for i in range(21) if i % 2 == 0]
print(res)

# 生成器表达式
res = (i for i in range(21) if i % 2 == 0)
print(res)

# 字典推导式
my_dict = {'hadoop': 1, 'spark': 2, 'kafka': 3}
res = {value: key for key, value in my_dict.items()}
print(res)