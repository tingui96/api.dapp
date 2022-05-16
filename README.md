```shell
export HLF_DAPP_CONFIG="/home/user_session_here/Downloads/api.hlf.evote.dapp/conf.linux.yaml"
export HLF_DAPP_JWT_SIGN_KEY="secrethatmaycontainch@r32lenght"
export HLF_DAPP_SISEC_CLIENT_ID="id_str"
export HLF_DAPP_SISEC_PW="the_password"
```
ex:

```shell
export HLF_DAPP_CONFIG="/home/user_session_here/Downloads/api.hlf.evote.dapp/conf.linux.yaml"
export HLF_DAPP_JWT_SIGN_KEY="45567f001601aacb761e13987cddc62ddd49c5b2"
```

Install Swagger api documentation running this command:

```shell
go install github.com/swaggo/swag/cmd/swag@v1.7.0
```

```shell
swag init --parseDependency --parseInternal --parseDepth 1 --md docs/md_endpoints
```

```shell
go build -o /home/portainer/go-path/bin/api.hlf.dapp
```


open the following url in the browser: http://192.168.49.133:7001/swagger/index.html
