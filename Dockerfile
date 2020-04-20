FROM golang:1.14.2 AS build

ENV TZ Asia/Shanghai

ENV GO111MODULE on

ENV GOPROXY https://goproxy.io

RUN echo 'Asia/Shanghai' >/etc/timezone

RUN apt-get update
RUN apt-get install expect -y
RUN apt-get install default-mysql-client -y

#TODO 系统依赖内容
