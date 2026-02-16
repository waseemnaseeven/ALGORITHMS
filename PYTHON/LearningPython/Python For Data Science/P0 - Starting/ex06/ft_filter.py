def myFunc(x):
  if x < 18:
    return False
  else:
    return True

def ft_filter(function, object):
    yield [item for item in object if function(item)]
    print("done")
    
    
def main():
    ages = [5, 12, 17, 18, 24, 32, 2, 63]
    adults = ft_filter(myFunc, ages)
    res = ft_filter(myFunc, [0,45,42,10])
    for x in adults:
        print(x)
    for y in res:
        print(y)
    print(adults)


if __name__ == "__main__":
    main()