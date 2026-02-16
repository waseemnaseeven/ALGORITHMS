def stars(L):
	for i in range(1, L + 1):
		for j in range(i):
			print("*", end="")
		print("")

stars(6)
