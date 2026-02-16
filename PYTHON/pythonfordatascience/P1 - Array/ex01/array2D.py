import numpy as np

def slice_me(family: list, start: int, end: int) -> list:

    if not isinstance(family, list):
        raise TypeError("Variable is not a list")
    
    shape = np.shape(family)
    print(f"My shape is : {shape}")
    slicer = family[start:end]
    shape = np.shape(slicer)
    print(f"My new shape is : {shape}")
    return slicer