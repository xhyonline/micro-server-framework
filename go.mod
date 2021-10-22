module github.com/xhyonline/micro-server-framework

go 1.16

require (
	github.com/BurntSushi/toml v0.3.1
	github.com/gogo/protobuf v1.3.2
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0
	github.com/stretchr/testify v1.7.0 // indirect
	github.com/xhyonline/xutil v0.1.20210904-0.20211022103614-a9bbbde7025c
	go.etcd.io/etcd v3.3.27+incompatible
	google.golang.org/grpc v1.38.0
	gorm.io/gorm v1.21.6
)

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
