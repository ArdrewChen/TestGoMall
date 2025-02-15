package main

import (
	"context"
	"fmt"

	"github.com/cloudwego/kitex/client"
	"github.com/yuefan-mo/studymall/demo/demo_thrift/kitex_gen/api"
	"github.com/yuefan-mo/studymall/demo/demo_thrift/kitex_gen/api/echo"
)

func main() {
	cli, err := echo.NewClient("demo_thrift", client.WithHostPorts("loaclhost:8888"))
	if err != nil{
		panic(nil)
	}

	res,err := cli.Echo(context.Background(), &api.Request{
		Message: "hello",
	})

	if err != nil{
		fmt.Println(err)
	}

	fmt.Printf("%v", res)
}