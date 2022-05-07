# loan-api
REST API for tracking user's balance and loans

## Installation & Run
```
# Download the project
$ go get github.com/mvrsss/loan-api

# Download Gin
$ go get github.com/gin-gonic/gin

# Download GORM
$ go get github.com/jinzhu/gorm

# Download JWT-go
$ go get github.com/dgrijalva/jwt-go

# Download Viper
$ go get github.com/spf13/viper

# Run
$ docker-compose -f docker-compose.yml up
```

## API
**/user/register**
* ```POST```: Register client

Request format:
```
curl --request POST \
  --url http://localhost:8080/user/register \
  --header 'Content-Type: application/json' \
  --data '{
    "name": "Alan",
    "surname": "Turing",
    "iin": "1234567890",
    "phonenumber": "+77774442167",
    "password": "enigma123",
    "address": "Euston Road, 96"
}'
```
Response format:
 ``` 
 { "token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiQWxhbiIsInN1cm5hbWUiOiJUdXJpbmciLCJpaW4iOiIxMjM0NTY3ODkwIiwicGhvbmVudW1iZXIiOiIrNzc3NzQ0NDIxNjciLCJleHAiOjE2NTE5NDg5NDF9.POhskhKz71umzuKq2i2eZ32Hoa1VRXaa5AjofdYmG6c"
 }
 ```
 
**/user/login**
* ```GET```: Login

Request format:
```
curl --request GET \
  --url http://localhost:8080/user/login \
  --header 'Content-Type: application/json' \
  --data '{
    "phonenumber": "+77774442167",
    "password": "enigma123"
}'
```
Response format:
```
{
"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiQWxhbiIsInN1cm5hbWUiOiJUdXJpbmciLCJpaW4iOiIxMjM0NTY3ODkwIiwicGhvbmVudW1iZXIiOiIrNzc3NzQ0NDIxNjciLCJleHAiOjE2NTE5NDg5NDF9.POhskhKz71umzuKq2i2eZ32Hoa1VRXaa5AjofdYmG6c",
"uid":"FPTNkDpNV9QkhDZRKK9hLj"
}
```

**/user/authorized/updatebalance/**
* ```POST```: Updates balance

Request format:
```
curl -X POST http://localhost:8080/user/authorized/updatebalance\
	-H "Content-Type: application/json"\
	-H "Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiQWxhbiIsInN1cm5hbWUiOiJUdXJpbmciLCJpaW4iOiIxMjM0NTY3ODkwIiwicGhvbmVudW1iZXIiOiIrNzc3NzQ0NDIxNjciLCJleHAiOjE2NTE5NDg5NDF9.POhskhKz71umzuKq2i2eZ32Hoa1VRXaa5AjofdYmG6c"\
	-d '{ 
        "phonenumber": "+77774442167",
        "password": "enigma123",
        "amount": 1000000
      }'
```
Response format:
```
{
    "balance": 1000000,
    "message": "user balance was successfully updated",
    "uid": "FPTNkDpNV9QkhDZRKK9hLj"
}
```

**/user/authorized/loanrequest/**
* ```POST```: Creates new loan requests

Request format:
```
curl -X POST http://localhost:8080/user/authorized/loanrequest\
	-H "Content-Type: application/json"\
	-H "Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiQWxhbiIsInN1cm5hbWUiOiJUdXJpbmciLCJpaW4iOiIxMjM0NTY3ODkwIiwicGhvbmVudW1iZXIiOiIrNzc3NzQ0NDIxNjciLCJleHAiOjE2NTE5NDg5NDF9.POhskhKz71umzuKq2i2eZ32Hoa1VRXaa5AjofdYmG6c"\
	-d '{
        "amount": 18000,
        "startdate": "2022-01-02T15:04:05Z",
        "enddate": "2022-02-02T15:04:05Z",
        "interestrate": 16.4,
        "clientid": "FPTNkDpNV9QkhDZRKK9hLj"
      }'
```
Response format:
```
{
    "message": "record saved"
}
```

**/user/authorized/getuserloan/**
* ```GET```: Gets all loan records of the given user

Request format:
```
curl -X POST http://localhost:8080/user/authorized/getuserloan\
	-H "Content-Type: application/json"\
	-H "Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiQWxhbiIsInN1cm5hbWUiOiJUdXJpbmciLCJpaW4iOiIxMjM0NTY3ODkwIiwicGhvbmVudW1iZXIiOiIrNzc3NzQ0NDIxNjciLCJleHAiOjE2NTE5NDg5NDF9.POhskhKz71umzuKq2i2eZ32Hoa1VRXaa5AjofdYmG6c"\
	-d '{
        "phonenumber": "+77774442167",
        "password": "enigma123",
        "clientid": "FPTNkDpNV9QkhDZRKK9hLj"
      }'
```
Response format:
```
{
    "Current balance": 982000,
    "Early Repayment Amount": 2.5688798554652217,
    "Next Payment Amount": 15000,
    "Next Payment date": "2022-08-08T15:04:05Z",
    "Payments Number": 1
}
```
