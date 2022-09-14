# Go final project

## Run with docker
```shell
docker-compose up
```
This command will execute all project docker file

This will provide : 
 - An api exposing its endpoint to create a player and a wallet
 - A mock api to simulate the creation of a waller on a blockchain which will be called by the other api
 - A postgres database to store created users and wallets

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

Example of body to provide :
```json
{
    "username" : "Hugo",
    "password" : "myComplexPassword",
    "pin_code" : "123456"
}
```

## Response 
### Success
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

### Error
This endpoint will respond with an error body like :
```json
{
    "error": "The error message"
}