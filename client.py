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

WATCH_DIRECTORY = "/home/groupf/Desktop/image"
SERVER_URL = "https://192.168.178.75/24:5000"
app = Flask(__name__)
processed_files = set()
lock = Lock()
HEADERS = {
    "X-API-KEY": "123456",
    "X-MAC-ADDRESS": "00:11:22:33"
    }

results_queue = Queue()

class FileEventHandler(FileSystemEventHandler):
    def on_created(self, event):
        if event.is_directory:
            return None
        else:
            filepath = event.src_path
            filename = os.path.basename(filepath)
            print(f"Detected file: {filename}")
            self.process_file(filepath)

    def process_file(self, filepath):
        with lock:
            if filepath in processed_files:
                return
            processed_files.add(filepath)
            print(processed_files)
        
        with open(filepath, 'rb') as file:
            response = requests.post(SERVER_URL, files={'image': file})
            print(response)
        
        with lock:
            if response.status_code == 200:
                print(f"File {filepath} processed successfully. Deleting file.")
                os.remove(filepath)
                results_queue.put('success')
            else:
                os.remove(filepath)
                print(f"Failed to process file {filepath}. Status code: {response.status_code}")
                results_queue.put('failed')

@app.route("/")
def index():
    return render_template("index.html")

@app.route("/take_photo")
def take_photo():
    camera = PiCamera()
    camera.start_preview()
    camera.capture('/home/groupf/Desktop/image/image.jpg')
    sleep(5)
    camera.stop_preview()
    
    
    return redirect(url_for("status"))

@app.route("/failed")
def failed():
    return render_template("failed.html")

@app.route("/success")
def success():
    return render_template("success.html")

@app.route("/status")
def status():
    if not results_queue.empty():
        result = results_queue.get()
        if result == 'success':
            return redirect(url_for('success'))
        elif result == 'failed':
            return redirect(url_for('failed'))
    return "No updates yet."

def monitor_folder(path):
    event_handler = FileEventHandler()
    observer = Observer()
    observer.schedule(event_handler, path, recursive=False)
    observer.start()
    try:
        while True:
            time.sleep(1)
    except KeyboardInterrupt:
        observer.stop()
    observer.join()

if __name__ == "__main__":
    monitor_thread = Thread(target=monitor_folder, args=(WATCH_DIRECTORY,))
    monitor_thread.daemon = True
    monitor_thread.start()
    
    app.run(debug=True, port=5000)