# Project 2 - Overlay Networking
Jaeik, Jessey, Ken and Joseph Z, deploy a multi-instance server and VNFs to a production environment on a cloud services platform of your choice. There must be communication across separate virtual machines managed by an SDN controller and service mesh, separate applications controlling and routing traffic between them. Infrastructure and application changes should be managed by an automated CI/CD pipeline.

## Recommended Tools
- [DockerHub](https://hub.docker.com/)
- [AWS](https://aws.amazon.com/)
- [Jenkins](https://jenkins.io/)
- [Terraform](https://www.terraform.io/)
- [Kubernetes](https://kubernetes.io/)

## Requirements
- [X] Documentation
- [X] Agile/Scrum Project Management [Jira](https://revaturepro.atlassian.net/secure/BrowseProjects.jspa)
- [X] Git Branching & Semantic Versioning
- [X] (Cloud) Production Environment
- [X] CI/CD Pipeline
- [X] Infrastructure as Code
- [X] Orchestration

## Presentation
- [X] 15-minute Demonstration
- [X] Presentation Slides
- [X] MVP Done

 
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
    