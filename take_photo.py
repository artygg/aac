from picamera import PiCamera
from time import sleep

# Initialize the PiCamera object
camera = PiCamera()

# Start the camera preview, which shows what the camera is seeing on the screen
camera.start_preview()

# Capture an image and save it to the specified file path
camera.capture('image/image.jpg')

# Keep the preview active for 5 seconds to give the user some feedback
sleep(5)

# Stop the camera preview
camera.stop_preview()