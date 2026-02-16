def positionEltListe(L, x):
    new = []
    for i in range(len(L)):
        if L[i] == x:
            new.append(i)
    if len(new) == 0:
        print("element", x, "is not in the list ", L)
    return new

print(positionEltListe([1,2,3,4,5,6,3], 7))