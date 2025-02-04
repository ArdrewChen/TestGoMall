package dal

import (
	"github.com/ArdrewChen/TestGoMall/demo/demo_thrift/biz/dal/mysql"
	"github.com/ArdrewChen/TestGoMall/demo/demo_thrift/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
