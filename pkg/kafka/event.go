/**
 * @Author pibing
 * @create 2020/11/15 12:10 PM
 */

package kafka

type ReportEvent struct {
	Id        string                 `json:"id"`               //事件id
	Topic     string                 `json:"topic"`            //主题
	Type      string                 `json:"type"`             //事件类型，格式：{模块名/报表名}_{事件名}
	Account   string                 `json:"account"`          //账号id  //可设置分区
	Time      string                 `json:"time"`             //事件时间，
	Body      map[string]interface{} `json:"body"`             //事件内容
	RetryTime string                 `json:"retryTime"`        //重试时间
	Num       int32                  `json:"num"`              //重试次数
}



type ReportFailEvent struct {
	Id        string                 `json:"id" `           //事件id
	Topic     string                 `json:"topic"`         //主题
	Type      string                 `json:"type"`          //事件类型
	Account   string                 `json:"account"`       //账号id  //可设置分区
	Time      string                 `json:"time"`          //事件时间，生成报表用的时间
	Body      map[string]interface{} `json:"body"`          //事件内容
	RetryTime string                 `json:"retryTime"`     //事件重试时间
	Num       int32                  `json:"num"`           //重试次数
}
