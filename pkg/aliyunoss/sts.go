
package aliyunoss

import (
	"crypto/hmac"
	"crypto/sha1"
	"crypto/tls"
	"encoding/base64"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

const OSS_TTL  =3600*4   //设置默认有效时间   此时间只能在角色最大会话时间范围内
type AliyunStsClient struct {
	ChildAccountKeyId  string
	ChildAccountSecret string
	RoleAcs            string
}

//认证信息
type credentials struct {
	AccessKeyId     string `json:"accessKeyId"`
	AccessKeySecret string `json:"accessKeySecret"`
	SecurityToken   string `json:"securityToken"`
	Expiration      string `json:"expiration"`
}
type assumedRoleUser struct {
	AssumedRoleId string `json:"assumedRoleId"`
	Arn           string `json:"arn"`
}
//获取临时token时候，阿里返回的数据结构体 ali := aliOssMsg{}
type aliOssMsg struct {
	RequestId       string          `json:"requestId"`
	AssumedRoleUser assumedRoleUser `json:"assumedRoleUser"`
	Credentials     credentials     `json:"credentials"`
	HostId          string          `json:"hostId"`
	Code            string          `json:"code"`
	Message         string          `json:"message"`
}



func NewStsClient(oss_key,oss_secret,oss_role_acs string) *AliyunStsClient {
	return &AliyunStsClient{
		ChildAccountKeyId:  oss_key,
		ChildAccountSecret: oss_secret,
		RoleAcs:            oss_role_acs,
	}
}
/*
sessionName       区别其它token
durationSeconds   有效时间
*/
func (cli *AliyunStsClient) GenerateSignatureUrl(sessionName, durationSeconds string) (string, error) {
	assumeUrl := "SignatureVersion=1.0"
	assumeUrl += "&Format=JSON"
	assumeUrl += "&Timestamp=" + url.QueryEscape(time.Now().UTC().Format("2006-01-02T15:04:05Z"))
	assumeUrl += "&RoleArn=" + url.QueryEscape(cli.RoleAcs)
	assumeUrl += "&RoleSessionName=" + sessionName
	assumeUrl += "&AccessKeyId=" + cli.ChildAccountKeyId
	assumeUrl += "&SignatureMethod=HMAC-SHA1"
	assumeUrl += "&Version=2015-04-01"
	assumeUrl += "&Action=AssumeRole"
	assumeUrl += "&SignatureNonce=" + UUID()

	//if durationSeconds == "" || durationSeconds == "0" {
	//	durationSeconds = OSS_TTL
	//}
	//ttl, _ := strconv.Atoi(durationSeconds)
	//if ttl < 900 || ttl > 3600 {
	//	durationSeconds = OSS_TTL
	//}
	durationSeconds =strconv.Itoa(OSS_TTL)
	assumeUrl += "&DurationSeconds=" + durationSeconds

	// 解析成V type
	signToString, err := url.ParseQuery(assumeUrl)
	if err != nil {
		return "", err
	}

	// URL顺序化
	result := signToString.Encode()

	// 拼接
	StringToSign := "GET" + "&" + "%2F" + "&" + url.QueryEscape(result)

	// HMAC
	hashSign := hmac.New(sha1.New, []byte(cli.ChildAccountSecret+"&"))
	hashSign.Write([]byte(StringToSign))

	// 生成signature
	signature := base64.StdEncoding.EncodeToString(hashSign.Sum(nil))

	// Url 添加signature
	assumeUrl = "https://sts.aliyuncs.com/?" + assumeUrl + "&Signature=" + url.QueryEscape(signature)

	return assumeUrl, nil
}

// 请求构造好的URL,获得授权信息
func (cli *AliyunStsClient) GetStsResponse(url string) ([]byte, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	return body, err
}
