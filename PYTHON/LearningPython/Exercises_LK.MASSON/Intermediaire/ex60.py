def calculPGCD(a,b):
    assert(a>0 and b>0)
    while b != 0:
        a,b = b, a%b
    return a

print(calculPGCD(3,5))
