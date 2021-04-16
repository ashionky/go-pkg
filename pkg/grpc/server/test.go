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
	list := make([]*protos.Respose_List, 0)

	for i := 0; i < 5; i++ {
		list = append(list, &protos.Respose_List{
			Id:   "id_" + strconv.Itoa(i),
			Name: "张三" + strconv.Itoa(i),
		})

	}
	resp.List = list

	return resp, nil
}
