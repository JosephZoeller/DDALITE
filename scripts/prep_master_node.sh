#!/bin/bash

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