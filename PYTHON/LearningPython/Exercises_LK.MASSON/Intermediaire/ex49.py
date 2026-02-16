def nbrValeurDict(d):
    # extraire liste des cles
    d_cle = list(d.keys())
    nbr = 0
    # parcourir toutes les cles
    for i in d_cle:
        longueur_val = len(d[i])
        nbr += longueur_val
    return (nbr)
print(nbrValeurDict({"a": [1,2,3], "b" : [3, "p"], "c" : [8]}))