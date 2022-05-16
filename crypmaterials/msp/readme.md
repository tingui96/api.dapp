> This folder holds the cert to authenticate the dapp operations 
> in the blockchain. The operations can be of 2 types: 
> - Normal TXs: the dapp we use the normal user cert
> - Admins OpsL the dapp we use the privilege admin cert
> 
> ❗ Its importan to clarify that this folder is only for development 
> purpose, the dapp config must define the actual identify folder to 
> use in runtime.


mkdir ~/dapp/crypmaterials/msp
cp -R crypto-config/organizations/org1.example.com/users/*  ~/dapp/crypmaterials/msp/

