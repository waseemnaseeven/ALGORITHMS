import numpy as np
from PIL import Image

def ft_load(path: str) -> list:

    dot = path.find('.')

    if dot < 0:
        raise ValueError("dot not found")

    format = path[dot:len(path)]
    print(f"the format is {format}")

    img = Image.open(path)
    rgb_im = img.convert('RGB')
    r, g, b = rgb_im.getpixel((1, 1))
    print(f"the shape of the image is : {r}, {g}, {b}")
    img_array = np.array(img)
    img.close()
    return img_array