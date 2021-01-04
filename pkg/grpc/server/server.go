package server

import (
	"fmt"
	"go-pkg/pkg/grpc/protos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"net"
	"time"
)

func GRPCServerInit(address string) (err error) {
	if address == "" {
		address = ":9999"
	}
	listen, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Printf("failed to listen: %v", err)
		return
	}
	keepParams := grpc.KeepaliveParams(keepalive.ServerParameters{
		MaxConnectionIdle:     time.Duration(60000000000),
		MaxConnectionAgeGrace: time.Duration(20000000000),
		Time:                  time.Duration(60000000000),
		Timeout:               time.Duration(20000000000),
		MaxConnectionAge:      time.Duration(7200000000000),
	})
	// 实例化grpc Server
	s := grpc.NewServer(keepParams)
	// 注册列表
	protos.RegisterTestServer(s, TestService) // 测试模块
	go func() {
		fmt.Println("gRPC start")
		if err := s.Serve(listen); err != nil {
		fmt.Println("gRPC.error: ", err)
			panic(err)
		}
		fmt.Println("gRPC end")
	}()
	fmt.Println("grpc init success port:",address)

	return
}
