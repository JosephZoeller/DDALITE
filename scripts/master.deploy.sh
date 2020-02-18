
export masterip=$(terraform output master_ip)
xfce4-terminal -e " ssh -i ../terraform/secrets/private.pem ubuntu@$masterip " &

