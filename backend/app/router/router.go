package router

import (
	"badmintonAppService/app/controller"
	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	playerRoutes := r.Group("v1/players")
	playerController := controller.NewPlayerController()
	playerRoutes.POST("page/list", playerController.PageList)
	playerRoutes.POST("player/list", playerController.PlayerList)
	return r
}
