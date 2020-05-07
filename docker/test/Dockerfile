FROM centos:7

LABEL author="issac_004"
LABEL version="go_1.14"

# 设置环境变量
ENV GOLANG_VERSION 1.14.1
ENV GOROOT /usr/local/src/go
ENV GOPATH /root/work/go_project
ENV PATH $PATH:$GOROOT/bin:$GOPATH/bin
ENV GOPROXY="https://goproxy.cn,direct"

# 进入目录 ，开始下载文件
WORKDIR /usr/local/src
#CMD echo "开始安装软件..."

# 安装依赖
RUN yum clear all; \
yum -y update ; \
yum -y install wget git ; \
wget https://studygolang.com/dl/golang/go$GOLANG_VERSION.linux-amd64.tar.gz && \
tar -xvf go$GOLANG_VERSION.linux-amd64.tar.gz && rm -rf go$GOLANG_VERSION.linux-amd64.tar.gz

# 查看go版本 ，创建gopath目录 ，给权限
RUN go version; \
	mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH";

WORKDIR $GOPATH/src
