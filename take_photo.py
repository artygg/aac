import cv2 as cv
from time import sleep

# Initialize the video capture object (0 for default camera, usually the built-in camera or first connected camera)
camera = cv.VideoCapture(0)

# Check if the camera is opened correctly
if not camera.isOpened():
    print("Error: Could not open camera.")
    exit()

# Start the camera preview
ret, frame = camera.read()

if ret:
    # Display the frame (this acts as a preview)
    cv.imshow('Camera Preview', frame)

    # Capture an image and save it to the specified file path
    cv.imwrite('image/image.jpg', frame)

    # Keep the preview active for 5 seconds to give the user some feedback
    cv.waitKey(5000)

    # Close the preview window
    cv.destroyAllWindows()
else:
    print("Error: Could not read frame from camera.")

# Release the camera
camera.release()
