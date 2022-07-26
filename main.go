package main

import (
	"ginessential/common"
	"ginessential/router"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
)

func main() {
	InitConfig()
	db := common.InitDB()
	defer db.Close()
	r := gin.Default()
	router.CollectRoute(r)
	port := viper.GetString("server.port")
	if port != "" {
		panic(r.Run(":" + port))
	}
	r.Run() // 默认8080端口
}

func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("application")       // 设置文件名
	viper.SetConfigType("yml")               // 设置文件类型
	viper.AddConfigPath(workDir + "/config") // 设置文件目录
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
