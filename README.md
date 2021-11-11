
# Golang 微服务 GRPC 标准框架(轻量级)

## 特性介绍

1. 可使用 `etcd` 集群或单节点作为注册中心
2. 客户端请求服务端自带负载均衡
3. 服务端启动后自动向 `etcd` 注册,默认每 10s 进行一次心跳续租
4. 自带优雅停止
5. panic recover
6. 服务端无需指定启动端口,他会自动生成,当然你也可以通过 WithPort() 自行设置
   当服务启动后,你可以通过命令 `netstat --nltp | grep 服务名` 进行查看
   启动的端口

## 使用说明

### 一、依赖说明

该框架依赖 `etcd` 因此你需要自行安装 `etcd`,并且在目录`configs/common/etcd.toml` 中进行配置

配置内容如下

```
[etcd]
host = "127.0.0.1"
port = 2379
```
你可以使用`docker`安装`etcd`单节点进行测试

**安装命令如下**

```
export HostIP=0.0.0.0
```
```
docker run -itd --rm \
  -p 2379:2379 \
  -p 2380:2380 \
  --volume=/etcd-data:/etcd-data \
  --name etcd quay.io/coreos/etcd:latest \
  /usr/local/bin/etcd \
  --data-dir=/etcd-data --name node1 \
  --initial-advertise-peer-urls http://${HostIP}:2380 --listen-peer-urls http://${HostIP}:2380 \
  --advertise-client-urls http://${HostIP}:2379 --listen-client-urls http://${HostIP}:2379 \
  --initial-cluster node1=http://${HostIP}:2380
```
进入容器测试命令
```
etcdctl --endpoints=http://${HostIP}:2379 member list
```


### 二、protobuf 使用说明
如果你是用来测试,以下内容可以略过,直接编译运行 `example` 目录下的 `server` 与 `client`即可。

命令:
```
go build server.go
go build client.go
```

但如若你需要进行开发,请自行安装 `protobuf` 编译工具

安装方法如下
#### 1. protoc 

下载地址:https://github.com/protocolbuffers/protobuf/releases

![](http://picture.xhyonline.com/imgs/2021/03/606322cadf88fd7b.png)

#### 2.安装 gogofaster 插件

`go get github.com/gogo/protobuf/protoc-gen-gogofaster`

#### 3.生成代码

编写在 proto 目录下编写代码,使用一下命令生成

`protoc --gogofaster_out=plugins=grpc:. ./目录地址/*`

并将生成的文件拖动至 `gen` 目录下,做到编码规范

### 结语

如果您喜欢该项目,请给个 star 吧

作者:兰陵美酒郁金香
