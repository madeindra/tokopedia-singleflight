
# Dummy Aps for Implementation Single Flight

This repo use for dummy aplication when u implementation using single flight golang and call endpoint on this repo. 




## Pre-instalation

U need install redis on your local and go lang with min version `1.16.x`


## Installation

1. clone this repo on your local
1. After clone, u need run this code on your terminal and folder target `go mod init` to initiate of import package. 
1. Run on your terminal `go mod tidy` to add module requirements and sum, this action will create new file `go.sum` and `go.mod`
1. Final step is u just type `go run main.go` and application already running. 
    
## API Reference

#### Get all items

```http
  GET /get_data
```
on this endpoint no need parameter request. But showing response parameters. 

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `error_code`      | `int` | Showing error code of http |
| `error_message`      | `string` | showing error message if any error|
| `result`      | `string` | result data|

example response: 
```json
{
  "error_code": 200,
  "error_message": "",
  "result": "get key_testing: Hello Data Founded"
}
```

