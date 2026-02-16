import time

start = time.time()
num = int(input("whats your integer: "))
print("num = {}".format(num))
for i in range(0, 11):
    print("{} x {} = {}".format(num, i, num * i))
print(time.time() - start)
