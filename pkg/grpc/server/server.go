package server

import (
	"context"
	"encoding/json"
	"fmt"
	grpcmiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"go-pkg/pkg/grpc/protos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
	"net"
	"strings"
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
	s := grpc.NewServer(
		grpc.UnaryInterceptor(grpcmiddleware.ChainUnaryServer(
			grpc.UnaryServerInterceptor(LogGrpcAccessMiddleware),   //日志打印
		)),
		keepParams)
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
	fmt.Println("grpc init success port:", address)

	return
}

func LogGrpcAccessMiddleware(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	ctx = context.WithValue(ctx, "requestTime", time.Now())
	//ctx = context.WithValue(ctx, "traceId", "11111")

	md, _ := metadata.FromIncomingContext(ctx)
	fmt.Println(md)

	caller:="NA"
	//if ok && len(md["caller"])>0{
	//	caller=md.Get("caller")[0]          //调用方服务名称
	//}
	if val, ok := md["caller"]; ok {
		caller = val[0]
	}
	traceId:="NA"
	if val, ok := md["x-b3-traceid"]; ok {
		traceId = val[0]
	}

	resp, _ = handler(ctx, req)
	q, _ := json.Marshal(req)
	p, _ := json.Marshal(resp)
	dur := time.Now().Sub(GetContextReqTime(ctx))
	path, _ := grpc.Method(ctx)
	ps, _ := peer.FromContext(ctx)
	ips:=strings.Split(ps.Addr.String(),":")
	ip:="NA"
	if len(ips)>0 {
		ip=ips[0]
	}
	//traceId := ctx.Value("traceId").(string)
	fmt.Println(ip,path,caller, string(q), string(p), dur, traceId)
	return
}

// 获取ctx中requestTime
func GetContextReqTime(ctx context.Context) time.Time {
	return ctx.Value("requestTime").(time.Time)
}
