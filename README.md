# go-kit学习

### demo1
- 简单web接口实现、基础功能封装、ctx使用
- 请求：curl -d '{"orderId": "112233"}' 127.0.0.1:8080/order/create
- 响应：{"code":0,"err":null}

### demo2
- 添加日志中间件、自定义中间件、ok-log日志追踪、Fail接口错误统一处理、进一步封装

### demo3
- 编写中...
- GRPC服务端Go语言编写、PHP客户端GRPC编写、以及GPRC封装
- PHP项目详见：

### demo4
- 编写中...
- 使用限流器、断路器、指标等功能，并使用命令行来控制功能

### demo5
- 编写中...
- 添加mysql和redis，以及配置文件操作

### demo6
- 编写中...
- 使用服务注册发现，使用Dockerfile和docker-compose命令来创建和管理容器

### 其他目录
- configs: 配置文件目录
- deployments: 部署配置目录，如：docker-compose.yml 
- internal: 微服务项目的共享和私有代码目录
- tools: 工具目录，可以从internal目录中导入代码使用
- 参考：https://github.com/golang-standards/project-layout