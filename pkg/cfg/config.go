/**
 * @Author pibing
 * @create 2020/11/14 1:34 PM
 */

package cfg

import (
"fmt"
"github.com/jinzhu/configor"
)

//配置文件对象
type Config struct {
	APPName string `default:"appname"`

	Server struct {
		Appmode    string `default:"dev"`
		Http_port  string `default:"80"`
	}

	Mysql struct {
		Host string      `default:"127.0.0.1:3306"`
		User      string `default:"root"`
		Password  string
		Dbname    string
		Charset   string `default:"utf8mb4"`
	}
	Mongodb struct{
		Host      string   `default:"127.0.0.1"`
		Port      string   `default:"27017"`
		User      string   `default:"root"`
		Password  string
		Dbname    string
		Poolsize   int     `default:"10"`
	}

	Redis struct {
		Host      string   `default:"127.0.0.1:6379"`
		Password  string   `default:""`
		Db         int     `default:"0"`    //库
		Max_active int     `default:"30"`   //最大连接数
	}
	Alioss struct {
	     Oss_key       string
	     Oss_secret    string
	     Oss_role_acs  string
		 Oss_bucket    string
    }
	Kafka struct{
		Url       string
		Partition int
		Topic     string
	}
	Es struct{
		Url      string
		User     string
		Password string
	}
}

//config 全局对象
var config = Config{}

func Initcfg(path string) error {

	//初始化config对象
	err := configor.Load(&config, path)
	if err != nil {
		fmt.Print("cfg init err!", err)
		return err
	}
	return nil
}

func GetConfig() *Config {
	return &config
}

