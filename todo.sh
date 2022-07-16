#!/bin/bash

#GO
echo "        ####################################################### "
echo "        #        generating RESTful API documentation         # "
echo "        ####################################################### "
swag init --parseDependency --parseInternal --parseDepth 1 --md docs/md_endpoints
go mod vendor
echo "        ####################################################### "
echo "        #        building the solution                        # "
echo "        ####################################################### "
go build
./api.dapp