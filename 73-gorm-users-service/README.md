## go get 

go get -u gorm.io/gorm
go get -u github.com/gin-gonic/gin

or 

go mod tidy


## Docker network

```docker network create demo-network
```

## docker postgres

```
docker run -d --name pg -p 5432:5432 --network demo-network -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -e POSTGRES_DB=usersdb postgres:16

docker run -d --name dbui -p 28080:8080 --network demo-network adminer

docker ps
```


# Create user login 
- give email and password
- if user logs int reponse as user successfully login

## JSON Webtoken

HEADER.PAYLOAD.SIGNATURE

- Header 


```json
{
"alg":"HS256", 
"typ":"JWT"
}
```
- There few other alg: RS256,ES256

- PAYLOAD

```json
{
"sub":"HS256",
"name":"Jiten",
"role":"admin",
"email":"jitenp@outlook.com"
}
```

- SIGNATURE

- HMACSHA256(base64UrlEncode(header)+"."+base64UrlEncode(payload)),secret

-- Task 

Write a middle ware to take the request data , method type and audit it in a file 

Every request that hits the router, it should be saved in a file in json format 
{
    "req_data":{},
    "headers":{}
}

// Create a new model for the product 
id, name, type, price, make, model 

You need to fill the actual create handler func

create all crud operations for user and also product 
GetuserByID, GetUserByUserName, UpdateUser, DeleteUser 
GetProductByID,GetProductByType,UpdateProduct,DeleteProduct,GetAllProducts
