#!/bin/bash
# The next portions are specific to the master node.

# Kubeadm prepares master node

kubeadm init



# To make kubectl work for non-root user
mkdir -p $HOME/.kube
   cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
   chown $(id -u):$(id -g) $HOME/.kube/config

# Start weave network add-on with kubernetes version passed in
# and also set as an env variable.

export kubever=$(kubectl version | base64 | tr -d '\n')
kubectl apply -f "https://cloud.weave.works/k8s/net?k8s-version=$(kubectl version | base64 | tr -d '\n')"

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

# IPtables setting
iptables -P FORWARD ACCEPT
sudo sysctl net.bridge.bridge-nf-call-iptables=1