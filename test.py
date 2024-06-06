from flask import Flask, request, jsonify
from flask_restful import Api, Resource, reqparse
import boto3
import requests
from flask_sqlalchemy import SQLAlchemy
import os

app = Flask(__name__)
api = Api(app)

region = 'us-east-2'  
rekognition_client = boto3.client('rekognition', region_name=region)
s3_client = boto3.client('s3', region_name=region)

expected_args = reqparse.RequestParser()
expected_args.add_argument("password", type=str, help="Password required", required=True)

# BASE = "http://127.0.0.1:5003/people/"
BASE2 = "http://172.20.10.3:8080/api/device/authorize"
BASE3 = "http://172.20.10.3:8080/api/device/attendance"
S3_BUCKET = 'new-bucket-to-recognise-faces'  

HEADERS = {
    'X-API-KEY': "123456",
    'X-MAC-ADDRESS': "00:11:22:33"
}

class Logic(Resource):
    def post(self):
        if 'image' not in request.files:
            return {"message": "No image file provided."}, 400

        # response = requests.get(BASE2, headers=HEADERS)
        # if response.status_code != 200:
        #     return {"message": "Failed to connect to the service."}, 404
        image_file = request.files['image']
        image_bytes = image_file.read()
        print("Waiting!")
        try:
            response = s3_client.list_objects_v2(Bucket=S3_BUCKET)
            if 'Contents' not in response:
                return {"message": "No reference images found in S3 bucket."}, 404
            
            reference_images = []
            for obj in response['Contents']:
                reference_image_key = obj['Key']
                reference_image = s3_client.get_object(Bucket=S3_BUCKET, Key=reference_image_key)
                reference_images.append((reference_image_key, reference_image['Body'].read()))

            for reference_image_key, reference_image_bytes in reference_images:
                rekognition_response = rekognition_client.compare_faces(
                    SourceImage={'Bytes': reference_image_bytes},
                    TargetImage={'Bytes': image_bytes},
                    SimilarityThreshold=65
                )
                
                if rekognition_response['FaceMatches']:
                    username = reference_image_key.split('.')[0]
                    print(username)
                    print("matched")
                    print(rekognition_response)
                    # post_response = requests.post(BASE3, json={"id": int(username)}, headers = HEADERS)
                    return "Success", 200
                # else:
                #     username = reference_image_key.split('.')[0]
                #     print("unmatched")
                #     print(username)
                #     print(rekognition_response)
                #     # post_response = requests.post(BASE3, json={"id": 0})
                #     return "Error", 404
            print("No matches")
            return "No matches found", 404
        except Exception as e:
            print(f"An error occurred: {e}")
            return {"message": "An internal error occurred."}, 500

if __name__ == "__main__":
    app.run(port=5001, debug=True, host="172.20.10.13") 
