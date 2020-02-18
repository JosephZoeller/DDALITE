##Local values pulled from var.json
locals {
  json_data = jsondecode(file("./var.json"))
  json_secrets= jsondecode(file("./secrets/creds.json"))
}
##Public IPs for Master EC2, 1 ip.
output "master_ip" {
  value = aws_instance.master.public_ip
  description = "The Private IP address of the server instance"
}

##AWS Login Settings and Setup
provider "aws" {
  access_key = local.json_secrets.access_key
  secret_key = local.json_secrets.secret_key
  region     = "us-east-2"
}
##SSH LOGIN KEYS
resource "aws_key_pair" "deployer" {
  key_name	  = "Key123"
  public_key	= file("./secrets/public.pub")
}
##EC2for MASTER 
resource "aws_instance" "master" {
  key_name = aws_key_pair.deployer.key_name
  ami           = local.json_data.MASTER_image_id
  instance_type = "t2.medium"
  security_groups = [aws_security_group.SSH.name]
  connection {
    user = "ubuntu"
    type = "ssh"
    private_key = file("./secrets/private.pem")
    host =  self.public_ip
    timeout = "4m"
  }
##Setup Directories for Master
  provisioner "remote-exec" {
    inline = [
    "mkdir -p terradir/secrets",
    "mkdir pods",
    "mkdir services",

    ]
  }
##Core Script
  provisioner "file" {
    source      = "../scripts/prep_core.sh"
    destination = "/tmp/prep_core.sh"
  }
##Master Script
  provisioner "file" {
    source      = "../scripts/prep_master_node.sh"
    destination = "/tmp/prep_master_node.sh"
  }
##Collider Service Yaml
  provisioner "file" {
    source      = "../collider-service.yaml"
    destination = "/home/ubuntu/services/collider-service.yaml"
  }
##Collider Pod Yaml
  provisioner "file" {
    source      = "../collider.yaml"
    destination = "/home/ubuntu/pods/collider.yaml"
  }
##Terraform Slave tf file
  provisioner "file" {
    source      = "../slave.tf"
    destination = "/home/ubuntu/terradir/slave.tf"
  }

##Place creds, and keys into secrets directory
  provisioner "file" {
    source      = "./secrets/"
    destination = "/home/ubuntu/terradir/secrets"
  }
##Place varraibles json into terradir directory
   provisioner "file" {
    source      = "var.json"
    destination = "/home/ubuntu/terradir/var.json"
  }

##Exicute Script
  provisioner "remote-exec" {
    inline = [
      "sudo /bin/bash /tmp/prep_core.sh",
      "sudo /bin/bash /tmp/prep_master_node.sh",
    ]
  } 
}
##Secuirty Group Allow SSH
resource "aws_security_group" "SSH" {
  name        = "allow_ssh"
  description = "Allow SSH traffic"
  ingress {
    from_port   = 0 
    to_port     = 0
    protocol =   "-1"

    cidr_blocks =  ["0.0.0.0/0"]
  }
  egress {
    from_port       = 0
    to_port         = 0
    protocol        = "-1"
    cidr_blocks     = ["0.0.0.0/0"]
  }
}

