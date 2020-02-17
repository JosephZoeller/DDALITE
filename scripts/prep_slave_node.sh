#!/bin/bash

# Kubeadm prepares slave node

# THIS IS SPECIFIC TO MASTER NODE RAN
kubeadm join 172.31.7.99:6443 --token 046ghg.uxse8grrrwubkjhc \
    --discovery-token-ca-cert-hash sha256:606a36e9da0ea62cf5eebd66406a71bb04be114f939648a3543a71715403a4ea


sudo su -
