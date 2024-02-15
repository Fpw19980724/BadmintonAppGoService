package main

import (
	"badmintonAppService/app/common"
	"badmintonAppService/app/router"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"os"
)

func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(workDir + "/app/config")

	err := viper.ReadInConfig()
	if err != nil {
		return
	}
}

func main() {
	// 初始化配置
	InitConfig()
	// 初始化数据库
	common.InitDB()

	// gin框架
	r := gin.Default()
	r = router.CollectRoute(r)

	// 定义test路由
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})

	// 启动Gin应用程序
	_ = r.Run()
}
