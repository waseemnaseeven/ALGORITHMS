import sys


def main():
    ac = len(sys.argv)

    assert ac == 2, f"Expected one arg, but got {ac - 1}"

    try:
        av = int(sys.argv[1])
    except ValueError:
        raise AssertionError("AssertionError: argument must be integers")

    if av % 2 == 0:
        print("I'm Even")
    elif av % 2 != 0:
        print("I'm Odd")


if __name__ == "__main__":
    main()
