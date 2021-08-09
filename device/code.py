import time
import usb_cdc
from adafruit_macropad import MacroPad

macropad = MacroPad()

class CommandParser:
  def __init__(self, input):
    self._input = input
  
  

while True:
    macropad.encoder_switch_debounced.update()
    if macropad.encoder_switch_debounced.pressed:
        usb_cdc.data.write("press!".encode())
        usb_cdc.data.flush()
    line = usb_cdc.data.readline()
    print(line.decode())
    time.sleep(0.4)


