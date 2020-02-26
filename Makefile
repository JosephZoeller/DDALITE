 .ONESHELL:

master:
	cd ./terraform
	sudo terraform init
	sudo terraform apply --auto-approve
	export masterip=$$(terraform output master_ip)
	ssh -i ./secrets/private.pem ubuntu@$$masterip

image:
	cd ./terraform/image
	sudo terraform init
	sudo terraform apply --auto-approve
	
destroy_master:
	cd ./terraform
	sudo terraform destroy --auto-approve

destroy_image:
	cd ./terraform/image
	sudo terraform destroy --auto-approve