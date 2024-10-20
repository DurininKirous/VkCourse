def findMaximum(array,n):
    maxElement = array[0]
    for i in range(1,n):
        if array[i] > maxElement :
            maxElement = array[i]
    return maxElement

n = int(input())
elements = input()
array = [int(x) for x in elements.split(' ')]
print(findMaximum(array,n))
