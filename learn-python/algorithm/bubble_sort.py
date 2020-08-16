# 冒泡排序
def bubble_sort(data):
    i = 0
    while i < len(data) - 1:
        j = 0
        while j < len(data) - 1 - i:
            if data[j] > data[j + 1]:
                data[j], data[j + 1] = data[j + 1], data[j]
            j += 1
        i += 1


if __name__ == "__main__":
    data = [2, 1, 6, 8, 3, 5, 9, 4, 7]
    bubble_sort(data)
    print(data)
