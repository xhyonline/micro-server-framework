
# 微服务 GRPC 标准框架未完

## 一、安装
### 1. protoc 

github上直接安装一个二进制

### 2.安装 gogofaster 插件

`go get github.com/gogo/protobuf/protoc-gen-gogofaster`

## 二、使用

编写 proto 并生成代码

`protoc --gogofaster_out=plugins=grpc:. ./dir/*`


## 三、docker 简易安装 ETCD

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
```
etcdctl --endpoints=http://${HostIP}:2379 member list
```