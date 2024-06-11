from picamera import PiCamera
from time import sleep

camera = PiCamera()

camera.start_preview()
camera.capture('image/image.jpg')
sleep(5)
camera.stop_preview()

