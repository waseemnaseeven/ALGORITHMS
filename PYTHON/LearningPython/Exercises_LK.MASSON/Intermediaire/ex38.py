def supprimerDoublons(L):
    L = list(set(L))
    L.sort()
    return L

print(supprimerDoublons([0,3,5,7,3,5,1,-1]))

