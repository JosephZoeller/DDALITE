variable "aws_access_key" {}
variable "aws_secret_key" {}

source "amazon-ebs" "packerex" {
    ami_name = "packer-example {{timestamp}}"
    instance_type = "t2.micro"
    region = "us-east-2"

    access_key = "${var.aws_access_key}"
    secret_key = "${var.aws_secret_key}"

    source_ami_filter {
      filters {
        virtualization-type = "hvm"
        name = "ubuntu/images/*ubuntu-xenial-16.04-amd64-server-*"
        root-device-type = "ebs"
      }
      owners = ["099720109477"]
      most_recent = true
    }

    communicator = "ssh"
    ssh_username = "ubuntu"
}

build {
  sources = [ "source.amazon-ebs.packerex" ]
}