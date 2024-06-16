# AWS face recognition 

For our attendance system we decided to use AWS servises, precisesly "AWS Rekognition". 
This service helps us identify people's personality on the photo which is sent by device and then take appropriate actions.

## Installation
### AWS 
1. Create an account in Amazon Web Services.
2. Create a user and give that user all needed permissions in order to have full access to the desired service.
3. Select your AWS region.
4. Create an s3 bucket to store images for the AWS Rekognition to process.

### Flask 
1. Install Flask (``pip3 install Flask``)
2. Set up the environment, by opening the terminal, selecting the desired folder and writing this command ``python3 -m venv {name of your environment}``

### boto3
1. Install boto3 (``pip3 install boto3``)

### Set up your AWS account
1. ``brew install awscli``
2. ``aws configure``
3. ```
    AWS Access Key ID [None]: YOUR_ACCESS_KEY_ID
    AWS Secret Access Key [None]: YOUR_SECRET_ACCESS_KEY
    Default region name [None]: YOUR_REGION (e.g., us-west-2)
    Default output format [None]: json
    ```
 
### Prerequisites
- AWS account
- boto3
- Flask 

## Usage
With the help of this part of the project you can build a system based on users` faces. 
You can manipulate output data in the way you want, for example, create a database of people's presence, or a diagram. 

