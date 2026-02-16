def numberOccurrence(L):
    count_dict = {}
    res = []
    
    # Comptabiliser les occurrences de chaque élément dans la liste
    for num in L:
        if num in count_dict:
            count_dict[num] += 1
        else:
            count_dict[num] = 1
    
    # Construire la liste de tuples
    for num in L:
        tuple_element = (num, count_dict[num])
        res.append(tuple_element)
    
    return res

print(numberOccurrence([-4, 8, -3, 2, 7, 9, -3, 8, 1]))

