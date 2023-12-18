## 一、启动完整版的Golang容器

### 1、编写Dockerfile文件
```dockerfile 
  # golang最新版本镜像
  FROM golang:latest 

  ENV GOPROXY https://goproxy.cn,direct

  # 镜像工作目录
  WORKDIR $GOPATH/src/github.com/moshushu/gin-blog
  # 拷贝当前路径下的所有文件到目标路径（$GOPATH/src/github.com/moshushu/gin-blog）
  COPY . $GOPATH/src/github.com/moshushu/gin-blog

  # 编译二进制文件
  RUN go build -o gin-blog .

  # 声明运行时容器提供服务端口
  EXPOSE 9000

  # 指定容器启动程序及参数
  ENTRYPOINT [ "./gin-blog" ]
```

### 2、关联mysql镜像
* 在config配置文件中将`host`修改为mysql:3306
  * mysql：指mysql镜像的名称


### 3、构建镜像
> 在项目根目录执行`docker build -t gin-blog`，编译docker镜像
* `-t`：指定镜像名称

### 4、启动go服务
> `docker run --link mysql:mysql -p 9000:9000 gin-blog`
* `--link`：将Golang容器和Mysql容器关联起来
* `mysql:mysql`：容器名称:容器别名


## 二、启动简洁版的Golang容器 -- Scratch

### 1、编写Dockerfile文件
```dockerfile 
  # golang最新版本镜像
  FROM scratch

  # 镜像工作目录
  WORKDIR $GOPATH/src/github.com/moshushu/gin-blog
  # 拷贝当前路径下的所有文件到目标路径（$GOPATH/src/github.com/moshushu/gin-blog）
  COPY . $GOPATH/src/github.com/moshushu/gin-blog

  # 声明运行时容器提供服务端口
  EXPOSE 9000

  # 指定容器启动程序及参数
  ENTRYPOINT [ "./gin-blog" ]
```

### 2、关联mysql镜像
> 与上面完整版一致不变

### 3、编译可执行文件
> `CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o gin-blog .`

### 4、构建镜像
> 与上面完整版一致不变

### 5、启动go服务
> 与上面完整版一致不变