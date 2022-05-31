Esta DApp esta configurada con los criptomateriales de la red `https://github.com/ic-matcom/test-network-optativo-nanobash`, recomendada para la tarea final.

```shell
# exportamos las variables de entorno (HLF_DAPP_CONFIG y HLF_DAPP_JWT_SIGN_KEY)
export HLF_DAPP_CONFIG="/home/user_session_here/Downloads/api.dapp/conf.linux.yaml"
export HLF_DAPP_JWT_SIGN_KEY="45567f001601aacb761e13987cddc62ddd49c5b2"
```

For install go plugin for generate swagger documentation, go to the root project folder and then run:

```shell
go install github.com/swaggo/swag/cmd/swag@v1.7.0
```

For to generate the swagger run:
```shell
swag init --parseDependency --parseInternal --parseDepth 1 --md docs/md_endpoints
```

Then run:
```shell
go mod tidy

go build -o /home/portainer/go-path/bin/api.dapp
```


open the following url in the browser: http://192.168.49.133:7001/swagger/index.html


copying the cryptomaterials
```shell
# creamos la estructura de carpetas dapp/crypmaterials/msp
mkdir -p ~/dapp/crypmaterials/msp
# dentro de msp copiamos los criptomateriales
cp -R ./crypmaterials/msp/*  ~/dapp/crypmaterials/msp/
```
