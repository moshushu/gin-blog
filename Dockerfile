# FROM golang:latest

# ENV GOPROXY https://goproxy.cn,direct

# WORKDIR $GOPATH/src/github.com/moshushu/gin-blog
# COPY . $GOPATH/src/github.com/moshushu/gin-blog

# RUN go build -o gin-blog .

# EXPOSE 9000
# ENTRYPOINT [ "./gin-blog" ]

# Scratch镜像，简洁、小巧，基本是个空镜像
FROM scratch

WORKDIR $GOPATH/src/github.com/moshushu/gin-blog
COPY . $GOPATH/src/github.com/moshushu/gin-blog

EXPOSE 9000
ENTRYPOINT [ "./gin-blog" ]

# CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o gin-blog .