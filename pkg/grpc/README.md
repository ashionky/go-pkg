# 如果没有安装protobuf，需要下载安装 protobuf 
cd protobuf-3.8.0
./configure
make
make check
make install

# 安装完成查看版本号，验证是否正确
protoc --version # 显示：libprotoc 3.8.0


vi ~/.bash_profile
# protobuf-3.8.0 是下载的protoc源码文件包，xxx注意路径
export PROTOBUF=/Users/xxx/protobuf-3.8.0
export PATH=$PROTOBUF/bin:$PATH


# 保存之后生效
source ~/.bash_profile
// 在 protoc-gen-go 同级目录下执行
// 只编译自己写的 .proto 文件，将 * 替换为自己编写的文件名
protoc -I protos/ protos/*.proto --go_out=plugins=grpc:protos
