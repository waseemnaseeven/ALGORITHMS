import os

path = os.getcwd()
name = os.path.basename(path)
ext = name.split(".")[-1]
print(ext)
