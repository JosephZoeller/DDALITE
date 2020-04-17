variable "aws_access_key" {}
variable "aws_secret_key" {}
variable "public_key" {}
variable "private_key" {}
variable "WORKER_image_id" {}

##Public IPs for WORKER EC2's, array of worker ips.
output "instance_ips" {
  value = "${aws_instance.worker.*.public_ip}"
  description = "The Private IP address of the server instance"
}
##AWS Login Settings and Setup
provider "aws" {
  access_key = "${var.aws_access_key}"
  secret_key = "${var.aws_secret_key}"
  region     = "us-east-2"
}
##SSH LOGIN KEYS
resource "aws_key_pair" "deployer" {
  key_name	  = "Key_slave"
  public_key	= "${var.public_key}"
}
##EC2's for SLAVES
resource "aws_instance" "worker" {
  count   = "${var.user_count}"
  key_name = "${aws_key_pair.deployer.key_name}"
  ami           = "${var.WORKER_image_id}"
  instance_type = "t2.micro"
  security_groups = [ "${aws_security_group.SSH_Worker.name}" ]
  connection {
    user = "ubuntu"
    type = "ssh"    
    private_key = "${var.private_key}"
    host =  "${self.public_ip}"
    timeout = "4m"
}
##Make file structure pods and services
    provisioner "remote-exec" {
    inline = [
      "mkdir pods",
      "mkdir services",      
    ]
  }
##Place scripts in tmp folder to run and delete after
   provisioner "file" {
    source      = "/home/ubuntu/terradir/prep_core.sh"
    destination = "/home/ubuntu/prep_core.sh"
  }
##Place scrpits in temp folder to run and delete after
  provisioner "file" {
    source      = "/home/ubuntu/terradir/prep_slave_node.sh"
    destination = "/home/ubuntu/prep_slave_node.sh"
  }
##Place varraibles json into terraform directory
   provisioner "file" {
    source      = "/home/ubuntu/terradir/mastertoken.json"
    destination = "/home/ubuntu/mastertoken.json"
  }

##Run scrpits for slave setup
  provisioner "remote-exec" {
    inline = [
      "sudo /bin/bash /home/ubuntu/prep_slave_node.sh",
    ]
  }   
}
##Allows SSH
resource "aws_security_group" "SSH_Worker" {
  name        = "allow_ssh_worker"
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
