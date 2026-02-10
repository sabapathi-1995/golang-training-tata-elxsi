# Golang 

- Create a directory 
- cd to the directory 

- To create a go project (binary project/library project)

```bash
go mod init demo
```

- Creates a go.mod file. go mod is a package manager, kind of a build tool

- To run Go application

```bash
go run main.go
go run .
```

- To run multiple Go files

```bash
go run main.go greet.go
go run .
```

- To build

```bash
go build main.go
# build with name 
go build -o demo main.go
# release build
go build -ldflags="-w -s" -o release_demo main.go
```

## Go environment variables

```  bash
go env
```

GOROOT, GOPATH,GOBIN,GOOS,GOARCH

## go cross compile options

```bash
 go tool dist list
 GOOS=windows GOARCH=amd64 go build -o slice_demo_win_amd.exe main.go
 ```
```bash
aix/ppc64
android/386
android/amd64
android/arm
android/arm64
darwin/amd64
darwin/arm64
dragonfly/amd64
freebsd/386
freebsd/amd64
freebsd/arm
freebsd/arm64
freebsd/riscv64
illumos/amd64
ios/amd64
ios/arm64
js/wasm
linux/386
linux/amd64
linux/arm
linux/arm64
linux/loong64
linux/mips
linux/mips64
linux/mips64le
linux/mipsle
linux/ppc64
linux/ppc64le
linux/riscv64
linux/s390x
netbsd/386
netbsd/amd64
netbsd/arm
netbsd/arm64
openbsd/386
openbsd/amd64
openbsd/arm
openbsd/arm64
openbsd/ppc64
openbsd/riscv64
plan9/386
plan9/amd64
plan9/arm
solaris/amd64
wasip1/wasm
windows/386
windows/amd64
windows/arm64
```

## Docker network

```docker network create demo-network
```

## docker postgres

```
docker run -d --name pg -p 5432:5432 --network demo-network -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -e POSTGRES_DB=usersdb postgres:16

docker run -d --name dbui -p 28080:8080 --network demo-network adminer

docker ps
```

### Golang Context Implementation

https://www.youtube.com/playlist?list=PLJE7PIP1qj_Rn9vq4V4jGJbj5KqEIWSUc