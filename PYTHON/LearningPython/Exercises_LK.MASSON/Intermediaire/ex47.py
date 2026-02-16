def verifMaj(phrase):
    bool = False
    for i in phrase:
        if i >= 'A' and i <= 'Z':
            bool = True 
    return (bool)

print(verifMaj("es légumes sont bon pour la santé"))