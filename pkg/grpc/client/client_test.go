/**
 * @Author pibing
 * @create 2021/3/8 12:56 PM
 */

package client

import (
	"context"
	"fmt"
	"go-pkg/pkg/grpc/protos"
	"go-pkg/pkg/grpc/server"
	"google.golang.org/grpc"
	"testing"
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

	GetList()
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
	for _, v := range r.List {
		fmt.Println("id==", v.Id, "name==", string(v.Name))
	}

	return
}
