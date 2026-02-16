def calculSomme(L):
    return sum(L)

print(calculSomme([3,2,6,9,-1,5]))

# OR

def calculateSum(L):
    ret = 0
    for i in L:
        ret = i + ret
    return ret
    
print(calculateSum([3,2,6,9,-1,5]))
