import sys

def main():
    upper = 0
    lower = 0
    punctuation_mark = 0
    space = 0
    digits = 0

    ac = len(sys.argv)

    assert ac == 2, f"Expected one arg, got {ac -1}"

    try:
        av = sys.argv[1]
    except AssertionError as msg:
        print(msg)

    for i in av:
        if i >= 'A' and i <= 'Z':
            upper += 1
        elif i >= 'a' and i <= 'z':
            lower += 1
        elif i >= '!' and i <= '/':
            punctuation_mark += 1
        elif i == ' ':
            space += 1
        elif i >= '0' and i <= '9':
            digits += 1
        else:
            continue

    print(f"The text contains {len(av)} characters")
    print(f"{upper} upper letters")
    print(f"{lower} lower letters")
    print(f"{punctuation_mark} punctuation marks")
    print(f"{space} spaces")
    print(f"{digits} digits")


if __name__ == "__main__":
    main()