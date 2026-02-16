import os
import qrcode

img = qrcode.make("https://www.youtube.com/watch?v=zL9HDw1kU28")

img.save("qr.png", "PNG")

os.system("open qr.png")
