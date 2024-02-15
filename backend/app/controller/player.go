package controller

import (
	"badmintonAppService/app/common"
	"badmintonAppService/app/models"
	"badmintonAppService/app/response"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
)

type IPlayerController interface {
	PageList(ctx *gin.Context)
	PlayerList(ctx *gin.Context)
}

type PlayerController struct {
	DB *gorm.DB
}

// NewPlayerController 返回操作对象
func NewPlayerController() IPlayerController {
	db := common.GetDB()
	return PlayerController{db}
}

func (p PlayerController) PageList(ctx *gin.Context) {
	// 获取分页参数
	pageNum, _ := strconv.Atoi(ctx.DefaultQuery("pageNum", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("pageSize", "10"))

	// 分页
	var players []models.Player
	p.DB.Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&players)

	// 记录的总条数
	var total int64
	p.DB.Model(models.Player{}).Count(&total)

	response.Success(ctx, gin.H{
		"data":  players,
		"total": total,
	}, "分页查询成功!")
}

func (p PlayerController) PlayerList(ctx *gin.Context) {
	// 获取参赛项目参数
	var data map[string]any
	err := ctx.ShouldBindJSON(&data)
	if err != nil {
		response.Fail(ctx, gin.H{}, "无效的请求参数")
	}
	event := data["event"].(string)
	var players []models.Player
	// 根据参赛项目选手分数倒排取，单打项目取前10，双打项目取前20
	if event == "MS" || event == "WS" {
		p.DB.Where("event = ?", event).Order("points desc").Limit(10).Find(&players)
	} else {
		p.DB.Where("event = ?", event).Order("points desc").Limit(20).Find(&players)
	}
	response.Success(ctx, gin.H{
		"data": players,
	}, "分项目查询成功！")
}
