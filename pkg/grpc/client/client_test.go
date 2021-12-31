/**
 * @Author pibing
 * @create 2021/3/8 12:56 PM
 */

package client

import (
	"context"
	"fmt"
	"github.com/satori/go.uuid"
	"go-pkg/pkg/grpc/protos"
	"go-pkg/pkg/grpc/server"
	"go-pkg/pkg/util"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"testing"
	"time"
)

//方便多个rpc客户端注册
var rpcTestclient struct {
	TestClient protos.TestClient
}

func init() {
	GetgRpcClient("")
}

func getconn(address string) (conn *grpc.ClientConn) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		panic(err)
		return
	}
	return
}

func GetgRpcClient(address string) {
	if address == "" {
		address = "127.0.0.1:9999"
	}
	if address != "" {
		conn := getconn(address)
		rpcTestclient.TestClient = protos.NewTestClient(conn)
	}
	if true {
		//其他客户端rpc注册
	}
}

//grpc 调用测试
func TestH(t *testing.T) {
	server.GRPCServerInit("") //grpc服务启动

	Getone("1")
	//Getone("2")
	//GetList()
}

func GetList() {
	//conn, err := GetgRpcClient("")
	//defer conn.Close()
	//// 初始化客户端
	//if err != nil {
	//	fmt.Println("gRPC client error :", err)
	//	return
	//}
	//c := protos.NewTestClient(conn)

	// 调用方法
	reqBody := new(protos.Request)
	//reqBody.Name = "TEST"

	r, err := rpcTestclient.TestClient.GetList(context.Background(), reqBody)
	if err != nil {
		fmt.Println("gRPC GetList err:", err)
		return
	}
	fmt.Println("rpc data:", r)
	for _, v := range r.Data {
		fmt.Println("id==", v.Id, "name==", string(v.Name))
	}

	return
}


func CreateCtxByBackground() context.Context {
	ctx:=context.Background()
	requestId := uuid.NewV4().String()
	spanid := util.Md5(requestId)
	traceid := util.Md5(spanid)
	newmd := metadata.MD{}

	newmd.Set( "x-b3-sampled", "1")
	newmd.Set( "x-b3-spanid", spanid[0:16])
	newmd.Set( "x-b3-traceid", traceid)
	newmd.Set( "x-request-id", requestId)
	newmd.Set( "caller", "mssiot_forum")

	ctx=metadata.NewOutgoingContext(ctx, newmd)
	ctx = context.WithValue(ctx, "traceId", traceid)
	ctx = context.WithValue(ctx, "requestTime", time.Now())
	return ctx
}

func Getone(id string) {

	// 调用方法
	reqBody := new(protos.Request)
	reqBody.Id = id
	newmd := metadata.MD{}
	newmd.Set("caller", "mssiot_user")

	ctx:=CreateCtxByBackground()
	//ctx = metadata.NewOutgoingContext(ctx, newmd)
	//ctx = metadata.AppendToOutgoingContext(ctx, "caller", "mssiot_user")

	_, err := rpcTestclient.TestClient.GetOne(ctx, reqBody)
	if err != nil {
		s, _ := status.FromError(err)
		fmt.Println("code:",s.Code())
		//fmt.Println("gRPC Getone err:", err)
		return
	}

	return
}
