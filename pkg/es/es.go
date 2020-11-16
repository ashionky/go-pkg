package es

import (
	"context"
	"fmt"
	"github.com/olivere/elastic"
	"go-pkg/pkg/cfg"
)

var client *elastic.Client
var ctx = context.Background()
var config  = cfg.GetConfig()


//初始化es客户端连接
func Init() error {
	var url = config.Es.Url
	var user = config.Es.User
	var password = config.Es.Password
	var err error
	client, err = elastic.NewClient(
		elastic.SetURL(url),
		elastic.SetBasicAuth(user, password),
		elastic.SetSniff(false),
	)
	if err != nil {
		return err
	}
	info, code, err := client.Ping(url).Do(ctx)
	if err != nil {
		return err
	}
	if code == 200 {
		fmt.Sprintf("connected to es: %s ,version: %s", info.ClusterName, info.Version.Number)
	}
	return nil
}

//获取es连接
func getClient() *elastic.Client {
	if client != nil {
		return client
	} else {
		Init()
		return client
	}
}

