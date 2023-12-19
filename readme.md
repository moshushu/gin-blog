## 项目功能新建顺序
  * 1、配置文件及配置文件读取
    * mode模式
    * 服务启动端口，及超时时间
    * 数据库连接配置
  * 2、编写错误码包 -- e包
    * code.go
    * msg.go
  * 3、编写工具包 -- util包
    * ...
  * 4、数据库连接 --> model-init
    * 数据库连接，得到db
    * 启动log
    * 设置空闲连接池中的最大连接数
    * 设置数据库打开连接的最大数量
  * 5、路由独立 -- routers
    * api 接口目录
    * v1  版本
  * 6、接口表单验证
    * beego/validation
  * 7、Jwt认证
    * jwt-go
    * token的生成（GenerateToken）和解析（ParseToken）
  * 8、Jwt认证中间件 -- middleware
    * 用于对接口做统一的token认证
  * 9、日志 -- 文件日志
    * 自定义log
      * 日志前缀
      * 日志级别
      * ...
    * os库使用
  * 10、优雅重启服务
    * 目的
      * 不关闭现有连接（正在运行中的程序）
      * 新的进程启动并替代旧进程
      * 新的进程接管新的连接、
      * 连接要随时响应用户的请求，当用户仍在请求旧进程时要保持连接，新用户应请求新进程，不可以出现拒绝请求的情况
    * github.com/fvbock/endless
      * endless 热更新是采取创建子进程后，将原进程退出的方式，这点不符合守护进程的要求
    * 如果你的Golang >= 1.8，也可以考虑使用 http.Server 的 Shutdown 方法
  * 11、Swagger接口文档
    * swag init （生成文档再启动服务）
    * https://github.com/swaggo/gin-swagger

  * 12、GORM Callbacks
    * 定制自己的Create、Update 和 Delete回调
    * 用于回写”created_on“、”modified_on“、“deleted_on”值

  * 13、重构第一步，对配置进行统管 -- 配置统管
    * 要避免多 init 的情况，尽量由程序把控初始化的先后顺序

  * 14、抽离File模块
    * 由于之前的logFile采用的并不是通用的File模块
    * 故，抽离出File模块做统一管理
    * 因此需将之前的logFiel对file的操作替换成File模块操作

  * 15、实现上传图片接口
    * 文件名称加密
    * 检查图片大小、图片后缀、获取图片完整访问URL、获取图片完整路径等
    * 实现图片上去的接口hander及路由
    * 实现http.FileServer，即上传图片访问的接口

## 第三方库
  * 1、com包
    * go常用的工具包
  * 2、ini包
    * 读取配置文件，比viper更轻量
  * 3、beego的Validation
    * 表单验证

## 项目模块对应功能
  * 1、