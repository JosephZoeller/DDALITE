#!/bin/bash

# This script will set up the node to be a master node.

echo Setting up Kubernetes Node...

# Update instance repos and install standard software.
sudo apt-get update

sudo su -

apt-get install \
    apt-transport-https \
    ca-certificates \
    curl \
    gnupg-agent \
    software-properties-common

# Add key for docker-ubuntu repo
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -

# Add repo to end of /etc/apt/source.list
add-apt-repository \
   "deb [arch=amd64] https://download.docker.com/linux/ubuntu \
   $(lsb_release -cs) \
   stable"

# Update new docker repo and install modern docker tools
apt-get update
apt-get install docker-ce docker-ce-cli containerd.io

# Update repo list with kubernetes tools and install the 3 universal tools
# all nodes are expected to have.
apt-get update && sudo apt-get install -y apt-transport-https curl
curl -s https://packages.cloud.google.com/apt/doc/apt-key.gpg | sudo apt-key add -
cat <<EOF | sudo tee /etc/apt/sources.list.d/kubernetes.list
deb https://apt.kubernetes.io/ kubernetes-xenial main
EOF
apt-get update
apt-get install -y kubelet kubeadm kubectl

# You can create a basic image from here if you care to.