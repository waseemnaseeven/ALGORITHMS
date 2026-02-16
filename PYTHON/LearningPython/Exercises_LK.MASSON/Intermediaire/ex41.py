def sommeSousListe(L, i, j):
	Lij = L[i:j+1]
	sum = 0
	for index in Lij:
		sum += index
	return sum
print(sommeSousListe([4,10,12,16,18],2,4))
