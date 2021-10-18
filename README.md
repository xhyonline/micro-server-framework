
# 微服务 GRPC 标准框架未完

## 一、安装
### 1. protoc 

github上直接安装一个二进制

### 2.安装 gogofaster 插件

`go get github.com/gogo/protobuf/protoc-gen-gogofaster`

## 二、使用

编写 proto 并生成代码

`protoc --gogofaster_out=plugins=grpc:. ./dir/*`

