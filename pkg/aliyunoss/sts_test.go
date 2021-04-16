/**
 * @Author pibing
 * @create 2020/11/14 1:00 PM
 */

package aliyunoss

import (
	"encoding/json"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"os"
	"testing"
)

//后台直接上传
func TestGoApiTemplate(t *testing.T) {
	region := "oss-cn-beijing.aliyuncs.com"
	accessKeyId := ""
	accessKeySecret := ""
	//stsToken :=""

	// 创建OSSClient实例。
	//option:=oss.SecurityToken(stsToken)   //sts访问时，需设置
	client, err := oss.New(region, accessKeyId, accessKeySecret)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}

	// 获取存储空间。
	bucket, err := client.Bucket("test")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	var file = "/Users/Desktop/WechatIMG23.png"
	// 上传本地文件。
	err = bucket.PutObjectFromFile("acs/qq.png", file)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}

}

//前端上传，到后端获取临时token信息
func TestGetToken(t *testing.T) {
	client := NewStsClient("", "", "") //配置中取值
	url, err := client.GenerateSignatureUrl("test", "")
	if err != nil {
		fmt.Print("GenerateSignatureUrl err")
		return
	}

	data, err := client.GetStsResponse(url)
	if err != nil {
		fmt.Print("GetStsResponse err")
		return
	}

	ali := aliOssMsg{}
	err = json.Unmarshal(data, &ali)
	if err != nil {
		fmt.Print("Unmarshal err")
		return
	}
	fmt.Print("token认证信息：", ali.Credentials)
}
