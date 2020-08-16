# 插入排序
def insert_sort(data):
    i = 1
    while i < len(data):
        j, elem = i - 1, data[i]
        while j >= 0 and data[j] > elem:
            data[j + 1] = data[j]
            j -= 1
        data[j + 1] = elem
        i += 1


if __name__ == "__main__":
    data = [2, 1, 6, 8, 3, 5, 9, 4, 7]
    insert_sort(data)
    print(data)
