/**
 * @Author pibing
 * @create 2021/1/4 10:15 AM
 */

package server

import (
	"context"
	"go-pkg/pkg/grpc/protos"
	"strconv"
)

type testService struct {
}

var TestService = testService{}

func (t testService) GetList(ctx context.Context, req *protos.Request) (*protos.Respose, error) {
	resp := new(protos.Respose)
	// todo 获取resp
    list :=make([]*protos.User,0)
	for i := 0; i < 5; i++ {
		list= append(list, &protos.User{
			Id:   "id_" + strconv.Itoa(i),
			Name: "张三" + strconv.Itoa(i),
		})

	}
	resp.Data = list
	resp.Code = 1
	resp.Msg = "ok"

	return resp, nil
}

func (t testService) GetOne(ctx context.Context, req *protos.Request) (*protos.Respose2, error) {
	resp := new(protos.Respose2)
	// todo 获取resp
    user:=new(protos.User)
	if req.Id=="1" {
		user.Id = "1"
		user.Name = "张1"
	}
	if req.Id=="2" {
		user.Id = "2"
		user.Name = "张2"
	}
	resp.Data = user
	resp.Code = 1
	resp.Msg = "ok2"
	return resp, nil
}
