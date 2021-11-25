import time
import usb_cdc
import json
from adafruit_macropad import MacroPad
import io

macropad = MacroPad()
usb_cdc.data.timeout = 0.1

COLORS = {
    "red": (255, 0, 0),
    "green": (0, 255, 0),
    "blue": (0, 0, 255),
    "yellow": (255, 255, 0),
    "cyan": (0, 255, 255),
    "magenta": (255, 0, 255),
    "white": (255, 255, 255),
    "off": (0, 0, 0),
}

class CommandParser:
  def __init__(self, input):
    self._input = input
  
class SetLight:
  def __init__(self, macropad):
    self._macropad = macropad
  def apply(self, command):
    pixel = command['pixel']
    color = COLORS[command['color']]
    self._macropad.pixels[pixel] = color

setlight = SetLight(macropad)
  

while True:
    key_event = macropad.keys.events.get()
    if key_event:
        usb_cdc.data.write("{}".format(key_event.key_number).encode())
        usb_cdc.data.flush()
    macropad.encoder_switch_debounced.update()
    if macropad.encoder_switch_debounced.pressed:
        print("Pressed")
        usb_cdc.data.write("press!".encode())
        usb_cdc.data.flush()
    line = usb_cdc.data.readline()
    if len(line) > 0:
      cmd = json.load(io.BytesIO(line))
      setlight.apply(cmd)
    time.sleep(0.4)


