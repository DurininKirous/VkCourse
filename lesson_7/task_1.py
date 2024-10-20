def findLastEvenContainer(array):
    stack = []
    for i in range(len(array)):
        if array[i]%2 == 0:
            stack.append(array[i])
    if stack:
        return stack.pop()
    else:
        return -1

n = int(input())
elements = input()
array = [int(x) for x in elements.split(' ')]
print(findLastEvenContainer(array))
