from enum import Enum, IntEnum


# 枚举
class VIP(Enum):
    YELLOW = 1
    YELLOW_ALIAS = 1
    GREEN = 2
    BLACK = 3
    RED = 4


# 数值类型枚举
class VIP1(IntEnum):
    YELLOW = 1
    GREEN = 2
    BLACK = 3
    RED = 4


if __name__ == '__main__':
    print(VIP.YELLOW)
    print(VIP.YELLOW.value)
    print(VIP.YELLOW.name)
    print(VIP['YELLOW'])

    for i in VIP:
        print(i.name, i.value)

    for i in VIP.__members__.items():
        print(i[0], i[1].value)
