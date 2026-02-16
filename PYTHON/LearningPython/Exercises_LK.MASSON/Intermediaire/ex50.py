def concatDict(d1, d2):
    ret = {}
    ret.update(d1)
    ret.update(d2)
    return ret

print(concatDict({"a" : 3, "b" : 6}, {"c" : 2, "d" : -1}))