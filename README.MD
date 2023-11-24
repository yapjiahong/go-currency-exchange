
## Init
- Download [go 1.20](https://go.dev/dl/)
- cd `<project>`
- initialize go `go mod init`
- Download dependency
  - `go mod download`


## Run
- `go run main.go`

## config port
- in `main.go`
```GO
srv := &http.Server{
    Addr:    ":8080",
    Handler: r,
}
``` 


## Endpoint

| Method | PATH      | Description       |
|--------|-----------|-------------------|
| GET    | /health   | health check      |
| GET    | /Currency | Exchange currency |

| Request Params | Type   | Required | Description              |
|----------------|--------|----------|--------------------------|
| Source         | string | o        | enum with: USD, TWD, JPY |
| Target         | string | o        | enum with: USD, TWD, JPY |
| Amount         | string | o        | example: $12,345         |

## example Request from browser
`http://localhost:8080/currency?source=USD&target=JPY&amount=$33`

- result
```json
{ 
 "msg":"success",
 "amount":"3,689.43"
}
```