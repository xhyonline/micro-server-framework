package main

import (
	"context"
	"fmt"
	"log"

	"github.com/xhyonline/micro-server-framework/gen/golang/basic"

	"github.com/xhyonline/micro-server-framework/gen/golang"

	"google.golang.org/grpc"
)

func main1() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	client := golang.NewRunnerClient(conn)
	result, err := client.Hello(context.Background(), &basic.Empty{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result.GetName())
}
