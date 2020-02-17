#!/bin/bash
# The next portions are specific to the master node.

# Kubeadm prepares master node


kubeadm init



# Make a directory with kubernetes configuration inside HOME.
# Give file super user permissions?
mkdir -p $HOME/.kube
   cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
   chown $(id -u):$(id -g) $HOME/.kube/config

# Check that what nodes are currently in cluster.
# Then check what pods are available across all namespaces.
kubectl get node
kubectl get pods --all-namespaces

# I have no idea what is going on here.

sysctl net.bridge.bridge-nf-call-iptables=1

# Start clouds.weave.works with kubernetes version passed in
# and also set as an env variable.

export kubever=$(kubectl version | base64 | tr -d '\n')
kubectl apply -f "https://cloud.weave.works/k8s/net?k8s-version=$(kubectl version | base64 | tr -d '\n')"

var10=$(sudo kubeadm token create --print-join-command)

arrvar10=(${var10// / })

masteripp=$(echo ${arrvar10[2]})

token=$(echo ${arrvar10[4]})

discovery_token=$(echo ${arrvar10[6]})

json_data="{
      \"masteripp\" : \"${masteripp}\",
      \"token\" : \"${token}\",
      \"discovery_token\": \"${discovery_token}\"
      }"

touch /home/mastertoken.json
chmod 777 /home/mastertoken.json 
echo $json_data | cat > /home/mastertoken.json

# Check for all nodes once more.
kubectl get nodes