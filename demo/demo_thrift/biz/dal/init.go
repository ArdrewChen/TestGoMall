package dal

import (
	"github.com/ArdrewChen/TestGoMall/biz/dal/mysql"
	"github.com/ArdrewChen/TestGoMall/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
