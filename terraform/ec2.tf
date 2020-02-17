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

##Public IPs for WORKER EC2's, array of worker ips.
output "instance_ips" {
  value = aws_instance.worker.*.public_ip
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
  ami           = local.json_data.WORKER_image_id
  instance_type = "t2.medium"
  security_groups = [aws_security_group.SSH.name]
  connection {
    user = "ubuntu"
    type = "ssh"
    private_key = file("./secrets/private.pem")
    host =  self.public_ip
    timeout = "4m"
  }
  provisioner "remote-exec" {
    inline = [
      "mkdir pods",
      "mkdir services",
    ]
  }
  # provisioner "file" {
  #   source      = "setup_master.sh"
  #   destination = "/tmp/setup_master.sh"
  # }
  provisioner "file" {
    source      = "../collider-service.yaml"
    destination = "/home/ubuntu/services/client-server.yaml"
  }
  provisioner "file" {
    source      = "../collider.yaml"
    destination = "/home/ubuntu/pods/html-server.yaml"
  }
  # provisioner "remote-exec" {
  #   inline = [
  #     "sudo /bin/bash /tmp/setup_master.sh",
  #   ]
  # }
}
##EC2's for WORKER
resource "aws_instance" "worker" {
  count   = 2
  key_name = aws_key_pair.deployer.key_name
  ami           = local.json_data.MASTER_image_id
  instance_type = "t2.micro"
  security_groups = [aws_security_group.SSH.name]
  connection {
    user = "ubuntu"
    type = "ssh"
    private_key = "${file("./secrets/private.pem")}"
    host =  self.public_ip
    timeout = "4m"
  }
}
##Secuirty Group Open everything
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

