##Builds Docker Image for blandwall. Tag is blandwall

docker build  -f ./build/blandwall.dockerfile -t=blandwall .

##RUN from root ex: ./scripts/bwsetup.sh