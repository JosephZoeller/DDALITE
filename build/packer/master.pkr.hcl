variable "aws_access_key" {}
variable "aws_secret_key" {}
variable "private_key_path" {}
variable "prep_core_path" {}

source "amazon-ebs" "packerex" {
    // AMI Configuration
    ami_name = "DDA_MASTER_IMAGE"// {{timestamp}}
    ami_description = "Provisioned master node for DDALITE"

    // Access Configuration
    region = "us-east-2"
    access_key = "${var.aws_access_key}"
    secret_key = "${var.aws_secret_key}"
    //ssh_private_key_file = "${var.private_key_path}"
    
    // Run Configuration
    instance_type = "t2.medium"
    source_ami_filter {
      filters {
        virtualization-type = "hvm"
        name = "ubuntu/images/*ubuntu-xenial-16.04-amd64-server-*"
        root-device-type = "ebs"
      }
      owners = ["099720109477"]
      most_recent = true
    }

    // Communicator Configuration
    communicator = "ssh"
    ssh_username = "ubuntu"
}

build {
  sources = [ "source.amazon-ebs.packerex" ]

  provisioner "shell" {
    script = "${var.prep_core_path}"
  }
}