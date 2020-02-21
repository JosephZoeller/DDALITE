#!/bin/bash

# Kubeadm prepares slave node

# THIS IS SPECIFIC TO SLAVE NODE RAN
kubeadm join $(sudo jq -r '.masteripp' mastertoken.json) --token $(sudo jq -r '.token' mastertoken.json) \
	--discovery-token-ca-cert-hash $(sudo jq -r '.discovery_token' mastertoken.json)


