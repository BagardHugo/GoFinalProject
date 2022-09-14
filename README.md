# Go final project

## Table of content

1. [Getting started](#Getting-started)
2. [Running](#Running)
3. [End point](#End-point)
4. [Testing](#End-point)
5. [Copyright](#End-point)


## Getting started

### Requirement

- Docker >= 20.10

### Instructions

Clone the repository

```sh
git clone https://github.com/BagardHugo/GoFinalProject.git
```

You  need to [install docker](https://docs.docker.com/get-docker/) first.
Docker is available on linux and MacOs and Windows ( using Docker Desktop).

## Running

### Run with docker compose
```shell
docker-compose up
```

This command will create 3 docker container
1. database
2. Api
3. Mock external Api

### Stop docker compose
If you need to stop the containers:

```sh
docker-compose down
```

## Call endpoint
The endpoint is exposed on the port 5001 :
```
http://localhost:5001/
```

The only method supported is POST

Header to provide :
```
Content-Type : application/json
```

Example of body :
```json
{
    "username" : "Hugo",
    "password" : "myComplexPassword",
    "pin_code" : "123456"
}
```

### Response 
#### Success
This endpoint will successfully respond with a body like :
```json
{
    "Id": 1,
    "username": "hugo",
    "password": "myComplexPassword",
    "pincode": "123456",
    "account": {
        "id": 1,
        "wallet_address": "829beafe-96c7-46dd-ae5f-0d35858dda8b",
        "currency_code": "ETH",
        "currency_balance": "0"
    }
}
```

#### Error
This endpoint will respond with an error body like :
```json
{
    "error": "The error message"
}
```

## Testing

For run unit test, execute command :

```sh
go test -v
```

## Benchmark

For run benchmark test, execute command :

```sh
go test -bench .
```


## Copyright
No copyright for this project :)