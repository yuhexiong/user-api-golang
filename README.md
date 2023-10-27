# User API

## Overview
- Lanuage: Golang
- Web Framework: Fiber
- DataBase: MongoDB

## ENV
copy .env.default and rename as .env
```
MONGO_URL=
DB_NAME=
API_PORT=
```

## Install Modules
```
go get -u && go mod tidy -v
```


## Run
```
go run main.go
```

## Test
```
go test -v ./test/.
```

## API

* `POST /user`
* `GET /user/:userId`
* `PUT /user/:userId`
* `PATCH /user/:userId/:status`
* `DELETE /user/:userId`
