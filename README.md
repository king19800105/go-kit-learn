# go-kit学习

### demo1
- 简单web接口实现，基础功能封装，ctx使用
- 请求：curl -d '{"orderId": "112233"}' 127.0.0.1:8080/order/create
- 响应：{"code":0,"err":null}

### 其他目录
- configs: 配置文件目录
- deployments: 部署配置目录，如：docker-compose.yml 
- internal: 微服务项目的共享和私有代码目录
- tools: 工具目录，可以从internal目录中导入代码使用