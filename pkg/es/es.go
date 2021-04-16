package es

import (
	"context"
	"fmt"
	"github.com/olivere/elastic"
	"go-pkg/pkg/cfg"
)

var client *elastic.Client
var ctx = context.Background()
var config = cfg.GetConfig()

//初始化es客户端连接
func Init() error {
	var url = config.Es.Url
	var user = config.Es.User
	var password = config.Es.Password
	fmt.Println(config.Es)
	var err error
	client, err = elastic.NewClient(
		elastic.SetURL(url),
		elastic.SetBasicAuth(user, password),
		elastic.SetSniff(false),
	)
	if err != nil {
		fmt.Println("NewClient err ", err)
		return err
	}
	info, code, err := client.Ping(url).Do(ctx)
	if err != nil {
		fmt.Println("Ping err ", err)
		return err
	}
	fmt.Println("code:", code)
	if code == 200 {
		fmt.Printf("connected to es: %s ,version: %s \n", info.ClusterName, info.Version.Number)
	}
	return nil
}

//获取es连接
func getClient() *elastic.Client {
	if client != nil {
		return client
	} else {
		err := Init()
		fmt.Println("init err ", err)
		return client
	}
}
