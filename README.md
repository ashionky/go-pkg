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

##整理各类工具方法,方便搭建基础框架
```json
配置中数据获取     pkg/cfg  
mysql、sqlite    pkg/db
redis：     pkg/redis
mongodb：   pkg/mongodb
消息队列：   kafka连接的初始化及生产、消费   pkg/kafka
excel:     表格的上传、下载     pkg/excel
阿里云oss:  后端直接上传或token获取返回前端上传   pkg/aliyunoss
日志:       输出到logs/log_20201116.log中，按日期区分可自行修改   pkg/log
http方法:       post、get使用     pkg/http
snowflake:     雪花id的生成       pkg/snowflake
跨域配置        pkg/middleware/cors.go
请求响应日志输出     pkg/middleware/logreq.go
```







