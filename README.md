
# go-gin-rest-api-clean-arch

A playground for exploring the Gin framework with clean architecture design

## Pre-requisites
- MySQL local client / use docker
- Go

## Setup

- Project directory and setup
```bash
  cd $GOPATH/src
  mkdir golang_rest_api && cd golang_rest_api
  touch server.go .env
  go mod init
```
- Install Gin
```bash
   go get -u github.com/gin-gonic/gin
```
- Install Go DotEnv
```bash
   go get github.com/joho/godotenv/cmd/godotenv
```
- Install Gorm and Gorm driver for mysql
```bash
   go get -u gorm.io/gorm
   go get -u gorm.io/driver/mysql
```

    
