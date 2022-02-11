FROM golang:1.16-alpine
MAINTAINER "xinmaozhu@foxmail.com"
# 设置go环境
ENV GOPROXY=https://goproxy.cn,direct \
    GO111MODULE=on
# 进入工作目录/auth
WORKDIR /work
ADD . .
# go mod 加载依赖包
# build 二进制文件
RUN go mod tidy \
    && CGO_ENABLED=0 GOOS=linux go build -o ./cmd/parse_server ./api/parse_server.go
# 暴露80端口
EXPOSE 80
# 程序入口点
CMD ["/work/cmd/parse_server"]


#FROM golang:1.16-alpine
#MAINTAINER "xinmaozhu@foxmail.com"
## 设置go环境
#ENV GOPROXY=https://goproxy.cn,direct \
#    GO111MODULE=on
## 进入工作目录/auth
#WORKDIR /work
#ADD . .
## go mod 加载依赖包
## build 二进制文件
#RUN go mod tidy \
#    && CGO_ENABLED=0 GOOS=linux go build -o ./cmd/sign_server ./api/sign_server.go
## 暴露8181端口
#EXPOSE 8181
## 程序入口点
#CMD ["/work/cmd/sign_server"]
