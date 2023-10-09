创建A、B、C三个服务
基于kubernetes部署，并通过NodePort开放FrontEnd服务
增加OTel系统，实现Trace追踪
集成DB，最好实现CDC系统
增加中间件，包括定时任务、nsq/kafka、redis

- 基于Docker构建镜像并启动
  - 构建镜像：`docker run -d -p 18080:18080  add-grpc-service:latest`
  - 启动服务(目前网络不通)：`docker-compose up -d`