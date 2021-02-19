#!/bin/bash
# tips:
# 1,protoc:下载linux版本的protoc编译器,拷贝到../bin目录下
# 2,linux下go get github.com/gogo/protobuf/protoc-gen-gofast,拷贝protoc-gen-gofast插件到本地(windows)GOPATH/bin目录下
# 3,本地(windows)下执行改脚本：Git Bash 客户端执行该脚本
if [[ -z `command -v protoc` ]];then
  alias protoc=../bin/protoc
fi

export PATH=$PWD/../bin/:$PATH

export GO_OUT_DIR='genereated_go'

for file in `ls ./proto/*`
do
    if test -f ${file}
    then
        protoc -I=. --go_out=${GO_OUT_DIR} ${file}
    fi
done

#cp -rf $GO_OUT_DIR/service/proto $GO_OUT_DIR
#rm -rf $GO_OUT_DIR/service