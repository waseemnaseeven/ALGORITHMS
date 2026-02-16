def VerifPresence(a, L):
    for i in L:
        if a == L[i]:
        	return True
        return False

print(VerifPresence(-1,[3,6,9,7,"abcr"]))
