import sys

def main():
    L = []

    try:
        ac = len(sys.argv)

        assert ac == 3, f"Need 2 args only, but we found {ac - 1 }"

        words = sys.argv[1]
        len_word = int(sys.argv[2])
    except AssertionError:
        raise AssertionError("the arguments are bad")

    L = words.split(' ')
    print(L)
    filter_function = lambda word: len_word <= len(word)
    L1 = [word for word in L if filter_function(word)]
    print(L1)

if __name__ == "__main__":
    main()