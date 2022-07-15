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

## Arquitectura

El proyecto trabaja en 3 capas: api, repositorio y servicio. Para extender las funcionalidades e integrarse con el chaincode
instalado, solo debe trabajar sobre los siguientes ficheros .go.

Capas | Descripción   | Documentación
--- | --- | ---
api/endpoints/end_blockchain_txs.go | En este .go se exponen los endpoint. Cada endpoint es una función go | https://www.iris-go.com/docs/#/?id=api-examples
service/svc_blockchain_txs.go | Es invocado por endpoint e interactua con la capa repositorio
repo/hlf/repo_blockchain.go | Interactua con el chaincode usando el SDK de go.


## Ejemplo 

```golang
// ReadAsset Get asset from the blockchain ledger. Contracts: mycc
// @Summary Get asset from the blockchain ledger.
// @description.markdown ReadAsset_Request
// @Tags Txs.mycc
// @Security ApiKeyAuth
// @Accept  json
// @Produce json
// @Param	Authorization	header	string	true 	"Insert access token" default(Bearer <Add access token here>)
// @Param	id				path	string	true	"ID"	Format(string)
// @Success 200 {object} byte "OK"
// @Failure 401 {object} dto.Problem "err.unauthorized"
// @Failure 400 {object} dto.Problem "err.processing_param"
// @Failure 502 {object} dto.Problem "err.bad_gateway"
// @Failure 504 {object} dto.Problem "err.network"
// @Router /txs/read_asset/{id} [get]
func (h HBlockchainTxs) ReadAsset(ctx iris.Context) {
	// checking the param
	id := ctx.Params().GetString("id")
	if id == "" {
		(*h.response).ResErr(&dto.Problem{Status: iris.StatusBadRequest, Title: schema.ErrProcParam, Detail: schema.ErrDetInvalidField}, &ctx)
		return
	}
	
	result, problem := (*h.service).ReadAssetSvc(id)
	if problem != nil {
		(*h.response).ResErr(problem, &ctx)
		return
	}

	(*h.response).ResOKWithData(result, &ctx)
}
```

- Los decoradores `@Summary`, `@description.markdown` agregan a la documentación del endpoint un resumen del endpoint y una descripción respectivamente. El decorador @description.markdown carga la descripción de un fichero en el camino `docs/md_endpoints/ReadAsset_Request.md`
- El decorador `@Tags`, agrupa el endpoint
- `@Param` es para definir parámetro de entrada al endpoint. En el ejemplo se establecen 2 parámetros (el token de acceso y un ID), ambos string y el primero es de tipo cabecera (header).
- `@Success` y `@Failure` define las posibles respuestas del endpoint, para ejecución de tipo satisfactoria y de tipo error respectivamente. Ese decorador va acompañado del código de respuesta, el tipo de dato y la respuesta.

En el endpoint `ReadAsset` obtenemos el ID pasado como parámetro e invocamos la función `ReadAssetSvc` de la capa de servicio. Luego la capa de servicio es la encargada de invocar la función correspondiente de la capa repositorio.
La capa de repositorio es la encargada de ejecutar el chaincode y la función (Tx) correspondiente empleando el SDK de fabric.

## TIPS

- Siempre que realice modificaciones en los endpoint `api/endpoints/end_blockchain_txs.go`  debes volver a generar la documentacion con el comando `swag`
- Si desea usar la Org2, debe modificar el fichero de configuracion `conf.linux_and_wsl.yaml` las variables (`MspId` y `CppPath`)
- Todos los criptomateriales estan en la carpeta `crypmaterials`

