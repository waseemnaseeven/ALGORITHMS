def filtrerMots(phrase, longueurMini):
    phrase = phrase.split(' ')
    ret = []
    for word in phrase:
        if len(word) >= longueurMini:
           ret = word
    print(ret)

filtrerMots("Quel est ton origine", 5)
