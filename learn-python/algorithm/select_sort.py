# 选择排序
def select_sort(data):
    i = 0
    while i < len(data):
        j, max = 0, 0
        while j < len(data) - i:
            if data[j] > data[max]:
                max = j
            j += 1
        data[max], data[len(data) - 1 - i] = data[len(data) - 1 - i], data[max]
        i += 1


if __name__ == "__main__":
    data = [2, 1, 6, 8, 3, 5, 9, 4, 7]
    select_sort(data)
    print(data)
