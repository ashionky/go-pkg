/**
 * @Author pibing
 * @create 2020/11/16 9:55 AM
 */

package mongodb

import (
	"fmt"
	"go-pkg/pkg/cfg"
	"testing"

)

type Posts struct {
	Id         string   `bson:"_id" json:"id"`
	Guid       string   `bson:"guid" json:"guid"`
	TagId      string   `bson:"tag_id" json:"tagId"`           // 标签id
	Source     int      `bson:"source" json:"source"`          // 1-个人 2-官方
	ViewNum    int      `bson:"view_num" json:"viewNum"`       // 浏览数
	CommentNum int      `bson:"comment_num" json:"commentNum"` // 评论数
	PraiseNum  int      `bson:"praise_num" json:"praiseNum"`   // 点赞量
	Status     int      `bson:"status" json:"status"`          // 1-显示 2-隐藏
	IsTop      int      `bson:"is_top" json:"isTop"`           // 1-置顶 2-不置顶
	IsFeature  int      `bson:"is_feature" json:"isFeature"`   // 1-置顶 2-不置顶
	Sort       int      `bson:"sort" json:"sort"`              // 排序
	PicURL     []string `bson:"pic_url" json:"picUrl"`         // 图片
	Body       []Body   `bson:"body" json:"body"`              // 帖子内容
	Site       string   `bson:"site" json:"site"`              // 站点
	CreateTime int      `bson:"create_time" json:"createTime"`
}

type ExternalContent struct {
	ExternalURL string `bson:"external_url"`
	Title       string `bson:"title"`
	ImgUrl      string `bson:"img_url"`
}

type Body struct {
	Language        string          `bson:"language"` // 语言
	Title           string          `bson:"title"`    // 标题
	Brief           string          `bson:"brief"`    // 摘要
	Content         string          `bson:"content"`
	ExternalContent ExternalContent `bson:"external_content"`
}

type T struct {
	Name string  `json:"name" bson:"name"`
	Age uint64 `json:"age" bson:"age"`
}
func TestMongoInit(t *testing.T) {
	var configFile = "../../conf/dev.yml"
	_ = cfg.Initcfg(configFile)
	config := cfg.GetConfig()
	//config.Mongodb.Host="merossdev.ejor5.mongodb.net"
	//config.Mongodb.Port="merossdev.ejor5.mongodb.net"
	//config.Mongodb.User="liaopibing"
	//config.Mongodb.Password="AQGVwjpLXAztNLNr"

	var url = "mongodb://" + config.Mongodb.User + ":" + config.Mongodb.Password + "@" + config.Mongodb.Host + ":" + config.Mongodb.Port + "/admin"

	url="mongodb+srv://liaopibing:AQGVwjpLXAztNLNr@merossdev.ejor5.mongodb.net/test?authSource=admin&replicaSet=atlas-zcqx9e-shard-0&readPreference=primary&appname=MongoDB%20Compass&ssl=true"
	err:=MongoInit(url, "test", config.Mongodb.Poolsize)
	fmt.Println(err)

	m:=T{}
	err = DbFindOne("mssiot_forum_merossbeta", "test", B{"name":"test"},&m)
	fmt.Println("==",m)
	fmt.Println("err",err)

}


