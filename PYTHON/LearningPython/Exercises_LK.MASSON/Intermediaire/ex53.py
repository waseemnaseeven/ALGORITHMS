def isThereAVoyel(sentence):
    voyels = ['a', 'e', 'i', 'o', 'u', 'y']
    for letter in voyels:
        if letter in sentence:
            return True
    return False

def presenceVoyelle(phrase):
    bool = False
    for i in phrase:
        if i == 'A' or i == 'a':
            bool = True
            return (bool)
        elif i == 'E' or i == 'e':
            bool = True
            return (bool)
        elif i == 'I' or i == 'i':
            bool = True
            return (bool)
        elif i == 'O' or i == 'o':
            bool = True
            return (bool)
        elif i == 'U' or i == 'u':
            bool = True
            return (bool)
        elif i == 'Y' or i == 'y':
            bool = True
            return (bool)
    return (bool)

print(presenceVoyelle("rbhpm"))
print(isThereAVoyel("rbhpm"))