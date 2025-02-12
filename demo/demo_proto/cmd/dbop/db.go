package main

import (
	// "fmt"

	"github.com/joho/godotenv"
	"github.com/yuefan-mo/studymall/demo/demo_proto/biz/dal"
	"github.com/yuefan-mo/studymall/demo/demo_proto/biz/dal/mysql"
	"github.com/yuefan-mo/studymall/demo/demo_proto/biz/model"
)

func main(){
	err := godotenv.Load()
	if err != nil{
		panic(err)
	}
	dal.Init()
	//mysql.DB.Create(&model.User{Email:"demo@example.com", Password:"zsctql"})
	// mysql.DB.Model(&model.User{}).Where("email=?", "demo@example.com").Update("password", "222222") // 修改数据
	// var row model.User
	// mysql.DB.Model(&model.User{}).Where("email=?", "demo@example.com").First(&row)//查询数据

	// fmt.Printf("row:%#v", row)
	// mysql.DB.Where("email=?", "demo@example.com").Delete(&model.User{}) // 删除，软删除 
	mysql.DB.Unscoped().Where("email=?", "demo@example.com").Delete(&model.User{}) //强制删除

}
