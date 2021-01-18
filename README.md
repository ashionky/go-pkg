go-pkg
### 使用说明
```
克隆代码
git clone https://github.com/ashion89/go-pkg.git

本地配置代理
export GOPROXY="https://goproxy.cn,direct"

依赖下载
go mod tidy 
go mod vendor

修改配置
conf/dev.yml

打包:
go build 

启动:
./go-pkg

```

## api服务基础架构；修改配置后可直接启动
## pkg目录下为整理封装的各类工具方法且已经测试，可直接使用，相关目录下有对应的_test文件
## 如使用中发现bug；欢迎大佬们指正；谢谢！
```json
配置中数据获取:                          --- pkg/cfg  
mysql:    gorm.io/gorm的使用            ---pkg/db 
          github.com/jinzhu/gorm的使用  ---pkg/zdb
redis:    go-redis的使用                ---pkg/go-redis 
          redigo的使用                  ---pkg/redis  
mongodb:                               ---pkg/mongodb
kafka:     kafka连接的初始化及生产、消费    ---pkg/kafka
es:       es连接初始化、数据存放、数据搜索   ---pkg/es
excel:     表格的上传、下载               ---pkg/excel
file:      文件上传到static下            ---pkg/file
grpc:      rpc服务创建/protoc安装        ---pkg/grpc
jwt:       对象生成token\解析token为对象  ---pkg/jwt_util
阿里云oss:  后端直接上传或token获取返回前端上传 ---pkg/aliyunoss
日志:       如输出到logs/log_20201116.log中,按日期区分可自行修改 ---pkg/log
http方法:       post、get使用            ---pkg/http
snowflake:     雪花id的生成              ---pkg/snowflake
跨域配置                                 ---pkg/middleware/cors.go
请求响应日志输出                          ---pkg/middleware/logreq.go
```







