创建A、B、C三个服务
基于kubernetes部署，并通过NodePort开放FrontEnd服务
增加OTel系统，实现Trace追踪
集成DB，最好实现CDC系统
增加中间件，包括定时任务、nsq/kafka、redis

- 基于Docker构建镜像并启动
  - 构建镜像：`docker run -d -p 18080:18080  add-grpc-service:latest`
  - 启动服务(目前网络不通)：`docker-compose up -d`
    - http://localhost:2379 替换成 http://my-etcd:2379
    - localhost:18080 替换成 :18080


如何创建etcd服务
```shell
docker run -d --name my-etcd -p 2379:2379 -p 2380:2380 quay.io/coreos/etcd /usr/local/bin/etcd --name my-etcd --advertise-client-urls http://0.0.0.0:2379 --listen-client-urls http://0.0.0.0:2379 --initial-advertise-peer-urls http://0.0.0.0:2380 --listen-peer-urls http://0.0.0.0:2380 --initial-cluster my-etcd=http://0.0.0.0:2380
```