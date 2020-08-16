# 希尔排序
def shell_sort(data):
    inc = len(data) // 2
    while inc > 0:
        i = 0
        while i < len(data):
            j, temp = i - inc, data[i]
            while j >= 0:
                if temp < data[j]:
                    data[j], data[j + inc] = data[j + inc], data[j]
                else:
                    break
                j -= inc
            i += 1
        inc //= 2


if __name__ == "__main__":
    data = [2, 1, 6, 8, 3, 5, 9, 4, 7]
    shell_sort(data)
    print(data)
