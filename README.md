# User API

### ENV
copy .env.default and rename as .env
```
MONGO_URL=
DB_NAME=
API_PORT=
```

### 更新 Modules
```
go get -u && go mod tidy -v
```


### 執行程式
```
go run main.go
```

### 執行測試
```
go test -v ./test/.
```

## API

### Request

* `POST /user`
* `GET /user/:userId`
* `PUT /user/:userId`
* `PATCH /user/:userId/:status`
* `DELETE /user/:userId`