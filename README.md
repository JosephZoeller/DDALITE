# Distributed Dictionary Attack (200106-uta-go Project-2)

## Overview
Jaeik, Jessey, Ken and Joseph Z's implementation of Project-2 for 200106-uta-go. This project is an automated provisioning system for instantiating multiple AWS EC2 machines as worker "collider" nodes, as well as building a Controller which distributes the workload evenly among the workers. The user can then request a decimal-formatted StrCode64 hash to be decrypted into its original string value.

Colliders are tasked with finding a StrCode64 hash collision for the user's inquiry. Upon recognizing a matching hash, the successful collider will return the hash - unhash pair to the Controller, who then returns the information to the client. The Collider Container is embedded with a literal dictionary for the Collider to reference as candidate, unhashed strings. StrCode64 is a particularly seeded cityhash64 function, and hashes can be generated using [BobDoleOwndU's QuickHash tool](https://github.com/BobDoleOwndU/QuickHash). More details on StrCode64 can be found [here](https://metalgearmodding.fandom.com/wiki/Hash_Wrangling). And, Although Colliders are designed to seek StrCode64 hashes, the logic can be swapped out with other hashing functions as necessity demands.

## Instructions before build.

AWS IAM Information and Credintials stroage
SETUP
- create secrets directory in project terraform directory
- create a creds.json file in the secrets directory
use this as a template

creds.json

    {

    "access_key" : "value of access key",
    "secret_key" : "value of secret key"

    }
- copy the access_key and secret_key from you aws IAM account into creds.json file. 
- place private key in secrets directory call it private.pem
- place public key in secrets directory call it public.pub

Setup Image ID befor Deployment

   
var.json Default values

    {
    "WORKER_image_id": "ami-0fc20dd1da406780b",
    "MASTER_image_id": "ami-0fc20dd1da406780b",
    "user_count": "3"
    }
Build the Worker/ Master Image ID

- change directory to the project root
- run the make command in terminal: make image
- terraform will build the core image with docker, kubernetes and all depedencies 
- once finished navagate to AWS website and go to running ec2 instance
- select the check on the newly spun ec2
- with it checked click the action button drop down to image then slide over to create image.
- name your image, list a discription, and create image 
- once the image has finished copy its ami number
- place the copy of the ami number in both master and worker_id of the var.json file
- now run the make command in terminal at project root: make destroy_image

## Setup Master Instance
- run make command in terminal: make master
- wait for the setup, estimated time is usually 2-4min
- once ssh connected follow the steps below to setup program:
    - sudo snap install go
    - go get -u -d github.com/200106-uta-go/JKJP2
    - cd go/src/github.com/200106-uta-go/JKJP2 
    - go build /src/revproxy
    - go build /src/sdnc

## Instructions on Connecting to Web Client
- Start the Web server and reverse proxy
    - sudo ./revproxy & disown
    - sudo ./sndc
- Web server is hosted and behind tls reverse proxy
- You need your Master's public ip to connect, can be optained from aws website on the master image dashboard, under public ip
    - type https://"your.masterip"
- Greeted with a client web page enter the Hash you want decifered and the priority of the build resources.
- wait for the collision detection to occur to decode the hash
    - Dashboard is simultaneously hosted for live feedback on the decodeing
        https://"your.masterup"?dash
    
