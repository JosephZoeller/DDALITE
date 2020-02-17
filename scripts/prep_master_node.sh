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
# The next portions are specific to the master node.

# Kubeadm prepares master node
sudu su -

kubeadm init

exit

# Make a directory with kubernetes configuration inside HOME.
# Give file super user permissions?
mkdir -p $HOME/.kube
  sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
  sudo chown $(id -u):$(id -g) $HOME/.kube/config

# Check that what nodes are currently in cluster.
# Then check what pods are available across all namespaces.
kubectl get node
kubectl get pods --all-namespaces

# I have no idea what is going on here.
sudo su -
sysctl net.bridge.bridge-nf-call-iptables=1

# Start clouds.weave.works with kubernetes version passed in
# and also set as an env variable.
exit 
export kubever=$(kubectl version | base64 | tr -d '\n')
kubectl apply -f "https://cloud.weave.works/k8s/net?k8s-version=$(kubectl version | base64 | tr -d '\n')"
exit

# Check for all nodes once more.
kubectl get nodes