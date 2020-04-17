 .ONESHELL:

master:
	cd ./terraform
	sudo terraform init
	sudo terraform apply --auto-approve
	export masterip=$$(terraform output master_ip)
	ssh -i ./secrets/private.pem ubuntu@$$masterip
	
destroy_master:
	cd ./terraform
	sudo terraform destroy --auto-approve

image:
	packer build \
		-var-file="build/secrets/creds.hcl" \
		-var 'private_key_path="build/secrets/private.pem"' \
		-var 'prep_core_path="build/scripts/prep_core.sh"' \
		build/packer/master.pkr.hcl
destroy_image:
	cd ./terraform/image
	sudo terraform destroy --auto-approve