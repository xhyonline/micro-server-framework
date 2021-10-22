package rpc

import (
	"context"

	"github.com/google/uuid"
	"github.com/xhyonline/micro-server-framework/gen/golang/basic"
	"github.com/xhyonline/micro-server-framework/gen/golang/hello"
)

type Service struct {
	Foo
}

type Foo struct {
}

var uid = uuid.NewString()

// Foo
func (s *Foo) Hello(context.Context, *basic.Empty) (*hello.Response, error) {
	return &hello.Response{Data: "你好世界,此消息来自机器:" + uid}, nil
}
