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

