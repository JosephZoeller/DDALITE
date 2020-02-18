#!/bin/bash

# Kubeadm prepares slave node

# THIS IS SPECIFIC TO MASTER NODE RAN
kubeadm join $(jq -r '.masteripp' mastertoken.json) --token $(jq -r '.token' mastertoken.json) \
	--discovery-token-ca-cert-hash $(jq -r '.discovery_token' mastertoken.json)


