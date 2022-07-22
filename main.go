package main

import (
	"ginessential/common"
	"ginessential/controller"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	db := common.InitDB()
	defer db.Close()
	r := gin.Default()
	r.POST("/api/auth/register", controller.Register)
	r.Run() // 默认8080端口
}
