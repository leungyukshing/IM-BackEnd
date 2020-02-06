##FROM alpine:latest
FROM golang:latest
#设置工作目录

RUN mkdir -p /go/src/github.com/backend

WORKDIR /go/src/github.com/backend
#将服务器的go工程代码加入到docker容器中 路径要修改
COPY . /go/src/github.com/backend

ENV ENV=product
#RUN go get github.com/astaxie/beego && go get github.com/beego/bee && go get github.com/go-sql-driver/mysql
##暴露端口
EXPOSE 9000

RUN pwd && cd /go/src/github.com/backend && pwd && ls
