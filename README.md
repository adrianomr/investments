# github.com/adrianomr/investments

### pre requisites

- make
- curl
- go 1.21 `make install-go`

### Running locally

1 - Install binaries.
> If this step fails, check if `GOPATH` is set correctly.
```shell
make install-binaries
```

2 - Defining initial config
```shell
make init
```

3 - Running app
```shell
make run
```

4 - Testing app
```shell
curl localhost:8080/public/demos 
```

### Running tests

```shell
make test
```

### Building app

```shell
make build
```

_Check if `.env` file exists on root folder._
 
```shell
./application
```

FINANCIAL DATA
https://github.com/Tpessia/dados-financeiros?tab=readme-ov-file

CDBS data

https://api.bcb.gov.br/dados/serie/bcdata.sgs.12/dados?formato=json