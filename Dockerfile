FROM golang:alpine AS builder

LABEL stage=gobuilder

#配置环境变量
ENV CGO_ENABLED 0
ENV GOOS linux
ENV GOPROXY https://goproxy.cn,direct


#容器启动时执行的命令会在该目录下执行
WORKDIR /app
ADD . /app


# 构建可执行文件
RUN go build go_ddns_namesilo


#最终运行docker的命令
ENTRYPOINT ["./go_ddns_namesilo"]

