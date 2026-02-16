class Person():
    # attributs
    def __init__(self, size, weight, age) -> None:
        self.size = size
        self.weight = weight
        self.age = age
    
    # methods
    def calculateIMC(self):
        res = self.weight / (self.size ** 2)
        print(f"{res:.2f}")
        return res

    def interpretationIMC(self):
        res = self.calculateIMC()
        if res <= 18.5:
            print("the person is very thin")
        elif res >= 30:
            print("the person is obese")
        else:
            print("the person is normal")

Waseem = Person(1.87,95,26)
Waseem.calculateIMC()
Waseem.interpretationIMC()