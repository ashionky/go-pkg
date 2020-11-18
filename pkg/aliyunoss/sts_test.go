/**
 * @Author pibing
 * @create 2020/11/14 1:00 PM
 */

package aliyunoss

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"os"
	"testing"
)

func TestGoApiTemplate(t *testing.T) {
	region := "oss-cn-beijing.aliyuncs.com"
	accessKeyId := ""
	accessKeySecret :=""
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
	var file="/Users/Desktop/WechatIMG23.png"
	// 上传本地文件。
	err = bucket.PutObjectFromFile("acs/qq.png", file)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}


}
