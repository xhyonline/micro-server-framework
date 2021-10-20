package rpc

import (
	"context"

	"github.com/xhyonline/micro-server-framework/gen/golang/basic"
	"github.com/xhyonline/micro-server-framework/gen/golang/user"
)

type Service struct {
	hello
}

type hello struct {
}

// hello
func (s *hello) Hello(context.Context, *basic.Empty) (*user.UserResponse, error) {
	return &user.UserResponse{Name: "小明"}, nil
}
