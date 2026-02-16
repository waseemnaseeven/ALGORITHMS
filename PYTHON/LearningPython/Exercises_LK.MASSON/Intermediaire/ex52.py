def diviseursMult(n, a, nbrSeuil):
    sum = []
    for i in range(1, nbrSeuil):
        if i % n == 0 and i % a != 0:
            sum.append(i)
    return sum
print(diviseursMult(5, 2, 100))
