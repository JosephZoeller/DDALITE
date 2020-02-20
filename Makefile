 .ONESHELL:

create master:
	cd ./terraform
	terraform init
	terraform apply --auto-approve
	export masterip=$$(terraform output master_ip)
	ssh -i ./secrets/private.pem ubuntu@$$masterip
