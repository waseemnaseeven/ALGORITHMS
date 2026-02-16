def sum(num):
   if num // 10 == 0:
       return num
   else:
       return num % 10 + sum(num // 10)

print(sum(149))
