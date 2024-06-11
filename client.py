import os
import time
import requests
from threading import Thread, Lock
from queue import Queue
from watchdog.observers import Observer
from watchdog.events import FileSystemEventHandler
from flask import Flask, render_template, redirect, url_for
from picamera import PiCamera
from time import sleep

# Directory to watch for new files
WATCH_DIRECTORY = "/home/groupf/Desktop/image"
# URL of the server to send the images
SERVER_URL = "https://192.168.178.75/24:5000"
# Initialize Flask app
app = Flask(__name__)
# Set of processed files to avoid re-processing
processed_files = set()
# Lock for thread-safe operations on processed_files
lock = Lock()
# Headers for the HTTP request
HEADERS = {
    "X-API-KEY": "123456",
    "X-MAC-ADDRESS": "00:11:22:33"
}

# Queue to store processing results
results_queue = Queue()


# Custom event handler for watchdog to handle file system events
class FileEventHandler(FileSystemEventHandler):
    def on_created(self, event):
        # Ignore directories
        if event.is_directory:
            return None
        else:
            # Process the created file
            filepath = event.src_path
            filename = os.path.basename(filepath)
            print(f"Detected file: {filename}")
            self.process_file(filepath)

    def process_file(self, filepath):
        # Lock to ensure thread-safe operation
        with lock:
            # Check if the file has already been processed
            if filepath in processed_files:
                return
            # Mark the file as processed
            processed_files.add(filepath)
            print(processed_files)

        # Send the file to the server
        with open(filepath, 'rb') as file:
            response = requests.post(SERVER_URL, files={'image': file})
            print(response)

        # Lock to ensure thread-safe operation
        with lock:
            # Check the server response and handle accordingly
            if response.status_code == 200:
                print(f"File {filepath} processed successfully. Deleting file.")
                os.remove(filepath)
                results_queue.put('success')
            else:
                os.remove(filepath)
                print(f"Failed to process file {filepath}. Status code: {response.status_code}")
                results_queue.put('failed')


# Flask route to serve the home page
@app.route("/")
def index():
    return render_template("index.html")


# Flask route to take a photo
@app.route("/take_photo")
def take_photo():
    # Initialize the PiCamera
    camera = PiCamera()
    # Start the camera preview
    camera.start_preview()
    # Capture an image and save it to the watched directory
    camera.capture('/home/groupf/Desktop/image/image.jpg')
    # Keep the preview for 5 seconds
    sleep(5)
    # Stop the camera preview
    camera.stop_preview()

    # Redirect to the status page
    return redirect(url_for("status"))


# Flask route to render the failure page
@app.route("/failed")
def failed():
    return render_template("failed.html")


# Flask route to render the success page
@app.route("/success")
def success():
    return render_template("success.html")


# Flask route to check the status of the last operation
@app.route("/status")
def status():
    # Check if there are any results in the queue
    if not results_queue.empty():
        result = results_queue.get()
        # Redirect based on the result
        if result == 'success':
            return redirect(url_for('success'))
        elif result == 'failed':
            return redirect(url_for('failed'))
    return "No updates yet."


# Function to monitor the specified folder for changes
def monitor_folder(path):
    # Create an event handler
    event_handler = FileEventHandler()
    # Initialize an observer
    observer = Observer()
    # Schedule the observer to watch the path
    observer.schedule(event_handler, path, recursive=False)
    # Start the observer
    observer.start()
    try:
        while True:
            # Keep the script running
            time.sleep(1)
    except KeyboardInterrupt:
        # Stop the observer on interruption
        observer.stop()
    observer.join()


# Main execution block
if __name__ == "__main__":
    # Start the folder monitoring in a separate thread
    monitor_thread = Thread(target=monitor_folder, args=(WATCH_DIRECTORY,))
    monitor_thread.daemon = True
    monitor_thread.start()

    # Run the Flask app
    app.run(debug=True, port=5000)