import numbers


class Group(object):
    def __init__(self, group_name, company_name, staffs):
        self.group_name = group_name
        self.company_name = company_name
        self.staffs = staffs

    # 实现序列协议，反转对象
    def __reversed__(self):
        self.staffs.reverse()

    # 实现切片操作
    def __getitem__(self, item):
        cls = type(self)
        if isinstance(item, slice):
            return cls(group_name=self.group_name,
                       company_name=self.company_name, staffs=self.staffs[item])
        elif isinstance(item, numbers.Integral):
            return cls(group_name=self.group_name,
                       company_name=self.company_name, staffs=[self.staffs[item]])

    # 实现长度判断操作
    def __len__(self):
        return len(self.staffs)

    # 实现可迭代性
    def __iter__(self):
        return iter(self.staffs)

    # 实现判断元素是否存在的特性
    def __contains__(self, item):
        if item in self.staffs:
            return True
        else:
            return False

    def __iadd__(self, other):
        if isinstance(other, self.__class__):
            return self.staffs.extend(other.staffs)

    def __str__(self):
        return ','.join(self.staffs)


if __name__ == '__main__':
    staffs = ['spark core', 'spark sql', 'spark streaming', 'spark mllib', 'spark graphx']
    group = Group(group_name='god', company_name='albert', staffs=staffs)

    # __reversed__测试
    reversed(group)
    print(group)

    # __getitem__测试
    print(group[2])
    print(group[: 2])

    # __len__测试
    print(len(group))

    # __iter__测试
    for elem in group:
        print(elem)

    # __contains__测试
    print('spark sql' in group)