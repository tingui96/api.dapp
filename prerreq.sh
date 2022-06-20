#!/bin/bash

echo "        ####################################################### "
echo "        #                   cloning the rest api              # "
echo "        ####################################################### "
cd $HOME
git clone https://github.com/ic-matcom/api.dapp.git
cd $HOME/api.dapp

#GO
echo "        ####################################################### "
echo "        #                   INSTALLING SWAG                   # "
echo "        ####################################################### "

go install github.com/swaggo/swag/cmd/swag@v1.7.0


#GO
echo "        ####################################################### "
echo "        #                   ENVIRONMENT                       # "
echo "        ####################################################### "

echo 'export HLF_DAPP_CONFIG="$HOME/api.dapp/conf.linux_and_wsl.yaml"' >> ~/.bash_profile
#echo 'export HLF_DAPP_CONFIG="/home/portainer/dapp/conf.linux.yaml"' >> ~/.bash_profile
#echo 'export HLF_DAPP_JWT_SIGN_KEY="45567f001601aacb761e13987cddc62ddd49c5b2"' >> ~/.bash_profile
source ~/.bash_profile

#cd $HOME
#mkdir dapp

#cp $HOME/api.dapp/conf.linux.yaml /home/portainer/dapp/
#cp $HOME/api.dapp/cpp.yaml /home/portainer/dapp/


echo "        ####################################################### "
echo "        #        generating RESTful API documentation         # "
echo "        ####################################################### "
cd $HOME/api.dapp
swag init --parseDependency --parseInternal --parseDepth 1 --md docs/md_endpoints
go mod tidy
go mod vendor
#go get


echo "        ####################################################### "
echo "        #        building the solution                        # "
echo "        ####################################################### "

go build
#go build -o $HOME/go-workspace/bin/api.dapp