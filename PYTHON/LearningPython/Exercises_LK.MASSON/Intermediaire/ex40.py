def maximum(L):
	ret = 0
	for i in L:
		if i >= ret:
			ret = i
	return ret
print(maximum([-9,2,4,1,8]))
