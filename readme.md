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

## 第三方库
  * 1、com包
    * go常用的工具包
  * 2、ini包
    * 读取配置文件，比viper更轻量
  * 3、beego的Validation
    * 表单验证

## 项目模块对应功能
  * 1、