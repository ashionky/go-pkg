/**
 * @Author pibing
 * @create 2021/1/4 10:15 AM
 */

package server

import (
	"context"
	"go-pkg/pkg/grpc/protos"
)

type testService struct {
}

var TestService = testService{

}

func (t testService) GetList(ctx context.Context, req *protos.Request) (*protos.Respose, error) {
	 resp := new(protos.Respose)
	  // todo 获取resp

	 return resp, nil
}
