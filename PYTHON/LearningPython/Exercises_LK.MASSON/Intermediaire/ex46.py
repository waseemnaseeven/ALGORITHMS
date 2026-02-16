def diviseur(n):
    n_diviseur = []
    # parcourir tous les éléments de 1 a n
    for div in range(1, n+1):
        if n % div == 0:
            n_diviseur.append(div)
    return n_diviseur

print(diviseur(9))