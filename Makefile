 .ONESHELL:

master:
	cd ./terraform
	terraform init
	terraform apply --auto-approve
	export masterip=$$(terraform output master_ip)
	ssh -i ./secrets/private.pem ubuntu@$$masterip

image:
	cd ./terraform/image
	terraform init
	terraform apply --auto-approve
	
destroy_master:
	cd ./terraform
	terraform destroy --auto-approve

destroy_image:
	cd ./terraform/image
	terraform destroy --auto-approve