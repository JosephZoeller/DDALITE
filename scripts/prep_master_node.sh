#!/bin/bash
# The next portions are specific to the master node

# Kubeadm initialization
# Flannel uses 10.244.0.0/16 as the pod network CIDR
kubeadm init --pod-network-cidr=10.244.0.0/16

# To make kubectl work for non-root user
mkdir -p $HOME/.kube
   cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
   chown $(id -u):$(id -g) $HOME/.kube/config

# Container Network Interface (CNI) Flannel installation
kubectl apply -f https://raw.githubusercontent.com/coreos/flannel/2140ac876ef134e0ed5af15c65e414cf26827915/Documentation/kube-flannel.yml

# IPtables configuration required by Container Network Interface (CNI), Flannel
sysctl net.bridge.bridge-nf-call-iptables=1

# Set environment variables
export kubever=$(kubectl version | base64 | tr -d '\n')

var1=$(sudo kubeadm token create --print-join-command)
arrvar1=(${var10// / })
masteripp=$(echo ${arrvar10[2]})
token=$(echo ${arrvar10[4]})
discovery_token=$(echo ${arrvar10[6]})
json_data="{
      var1: ${var1},
      arrvar1: ${arrvar1},
      masteripp: ${masteripp},
      token: ${token},
      discovery_token: ${discovery_token}
      }"
touch ../terraform/mastertoken.json
cat $json_data > ../terraform/mastertoken.json