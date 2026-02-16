L1 = [9,8,7,14,3,2,"a","p","bonjour","b"]
L2 = ["b",1,9.2,6,3,9,"p"]
L3 = []
for i in range(len(L1)):
    for j in range(len(L2)):
        if L1[i] == L2[j]:
            L3.append(L1[i])
print(L3)
