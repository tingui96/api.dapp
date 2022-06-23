# DAPP
Esta DApp está configurada con los criptomateriales de la red `https://github.com/ic-matcom/test-network-optativo-nanobash`, recomendada para la tarea final.

Antes recuerde modificar en el fichero de configuracion `conf.linux_and_wsl.yaml` la variable `ApiDocIp` y colocar el IP de la MV.
Si lo esta ejecutando en el HOST puede usar `127.0.0.1`

## Instalacion Automatica (solo ha sido probado en la MV del optativo)

Ejecute el script `prerreq.sh` que se encuentra en la raiz de la dapp. Este script realiza las siguientes tareas:

- Instala SWAG
- Exporta las Variables de entorno necesarias
- Genera la doc de swagger con la herramienta swag
- Compila la solución (DAPP)

```shell
./prerreq.sh
```

```shell
source ~/.bash_profile
```

## Instalacion Manual
Procedimiento manual para la instalación y preparación del entorno de la DAPP. 

La variable de entorno `HLF_DAPP_CONFIG` se configura con el camino absoluto al fichero de configuración `conf.linux_and_wsl.yaml`, por ejemplo:

```shell
# exportamos la variable de entorno (HLF_DAPP_CONFIG)
export HLF_DAPP_CONFIG="/home/portainer/api.dapp/conf.linux_and_wsl.yaml"
```

### Para instalar el módulo swag de go para generar la documentación swagger, debe moverse a la carpeta raíz del proyecto y ejecutar:

```shell
go install github.com/swaggo/swag/cmd/swag@v1.7.0
```

### Para generar la documentación, debe ejecutar:
```shell
swag init --parseDependency --parseInternal --parseDepth 1 --md ./docs/md_endpoints
```

### Luego debe ejecutar:
```shell
go mod vendor

go build -o api.dapp
```

## Ejecutar DAPP 
Para ejecutar la dapp solo ejecute el binario compilado:
```shell
./api.dapp
```

open the following url in the browser:

```shell
# usar esta URL SI esta usando la MV (remplazar `IP_de_la_MV` por el IP de la MV)
http://IP_de_la_MV:7001/swagger/index.html

# usar esta URL SI esta ejecutando la dapp en el HOST
http://127.0.0.1:7001/swagger/index.html
```


## TIPS

- Siempre que realice modificaciones en los endpoint `api/endpoints/end_blockchain_txs.go`  debes volver a generar la documentacion con el comando `swag`
- Si desea usar la Org2, debe modificar el fichero de configuracion `conf.linux_and_wsl.yaml` las variables (`MspId` y `CppPath`)
- Todos los criptomateriales estan en la carpeta `crypmaterials`
