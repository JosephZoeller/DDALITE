##Local values pulled from var.json
locals {
  json_secrets= jsondecode(file("../secrets/creds.json"))
}
#AWS Login Settings and Setup
provider "aws" {
  access_key = local.json_secrets.access_key
  secret_key = local.json_secrets.secret_key
  region     = "us-east-2"
}
##SSH LOGIN KEYS
resource "aws_key_pair" "deployer" {
  key_name      = "Key_Image"
  public_key    = file("../secrets/public.pub")
}
##EC2for MASTER 
resource "aws_instance" "image" {
  key_name = aws_key_pair.deployer.key_name
  ami           = "ami-0fc20dd1da406780b"
  security_groups = [aws_security_group.SSH.name]
  instance_type = "t2.medium"
  connection {
    user = "ubuntu"
    type = "ssh"
    private_key = file("../secrets/private.pem")
    host =  self.public_ip
    timeout = "4m"
  }

    ##Core Script
  provisioner "file" {
    source      = "../../scripts/prep_core.sh"
    destination = "/home/ubuntu/prep_core.sh"
  }

    ##Exicute Script
  provisioner "remote-exec" {
    inline = [
      "sudo /bin/bash /home/ubuntu/prep_core.sh",
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