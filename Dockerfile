FROM golang:1.11.2
ENV GOPATH /go
RUN go get github.com/beego/bee
RUN go get github.com/astaxie/beego
RUN go get github.com/go-sql-driver/mysql
ENV PATH $PATH:$GOPATH/bin
# ADD github.com /go/src/github.com

RUN mkdir -p /go/src/myblog

# 将工作目录切换到 /go/src/myblog 下
WORKDIR /go/src/myblog

# 将微服务的服务端可运行文件拷贝到 /app 下
COPY . /go/src/myblog
EXPOSE 8000
CMD ["build.sh"]
RUN bee run

