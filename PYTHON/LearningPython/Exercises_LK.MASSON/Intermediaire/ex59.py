def unionList(L1, L2, L3):
    new = []
    new_list = L1 + L2 + L3
    new_list.sort()
    new = set(new_list)
    return new

print(unionList([3,6,9,3],[1,0,3],[12,6,0]))