rangeserver=$(terraform output -json instance_ips | jq  .[] )
j=0
#range over our ips in the ip array
for i in $rangeserver; do
export serverip=$(terraform output -json instance_ips | jq .[$j])
xfce4-terminal -e " ssh -i ../terraform/secrets/private.pem ubuntu@$serverip " 
#increment j
((j=j+1))
done



