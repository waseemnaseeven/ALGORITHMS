def minimum(L):
    sum = 0
    for i in L:
        if (i < sum):
            sum = i
    return sum

print(minimum([-9,2,4,1,8]))
