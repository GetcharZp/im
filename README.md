# IM

> IM(即时通讯)
> 
> 往期视频：https://www.bilibili.com/video/BV1rJ411p7Mi/

## 技术栈
语言：Golang 数据库：MongoDB 框架：GIN 协议：Websocket

## 核心包
https://github.com/gorilla/websocket

## 扩展安装
```shell
go get -u github.com/gin-gonic/gin
go get github.com/gorilla/websocket
```

## Docker 安装 mongoDB
```shell
docker run -d --network some-network --name some-mongo \
-e MONGO_INITDB_ROOT_USERNAME=admin \
-e MONGO_INITDB_ROOT_PASSWORD=admin \
-p 27017:27017 \
mongo
```
