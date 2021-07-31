import time
import usb_cdc
from adafruit_macropad import MacroPad

macropad = MacroPad()

while True:
    usb_cdc.data.read(1)
    print("got it!")
    time.sleep(0.4)
