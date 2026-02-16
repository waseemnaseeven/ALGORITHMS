def readFile(path):
    file = open(path, "r")
    print(file)
    content = file.read()
    file.close()
    return content

print()