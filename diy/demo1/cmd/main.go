package main

import (
	"fmt"
	"github.com/golang/protobuf/ptypes"
	"grpc_practice/diy/demo1/genereated_go/proto/entity"
)

func main() {
	marshal := &entity.TestAny{
		Id:            1,
		Title:         "标题",
		Content:       "内容",
	}
	any, err := ptypes.MarshalAny(marshal)
	fmt.Println(any, err) // [type.googleapis.com/rpc.TestAny]:{Id:1 Title:"标题" Content:"内容"} <nil>

	msg := &entity.Response{
		Code: 0,
		Msg:  "success",
		Data: any,
	}
	fmt.Println(msg) // Msg:"success" data:{[type.googleapis.com/rpc.TestAny]:{Id:1 Title:"标题" Content:"内容"}}

	unmarshal := &entity.TestAny{}
	err = ptypes.UnmarshalAny(msg.Data, unmarshal)
	fmt.Println(unmarshal, err) // Id:1 Title:"标题" Content:"内容" <nil>
}
