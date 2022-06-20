#!/bin/bash

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
source ~/.bash_profile


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

go build -o api.dapp