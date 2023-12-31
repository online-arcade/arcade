package api

import (
	"arcade/types"
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/v3/pkg/curd"
)

// @Summary 查询签到记录数量
// @Schemes
// @Description 查询签到记录数量
// @Tags signin
// @Param search body ParamSearch true "查询参数"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyData[int64] 返回签到记录数量
// @Router/signin/count [post]
func noopSignInCount() {}

// @Summary 查询签到记录
// @Schemes
// @Description 查询签到记录
// @Tags signin
// @Param search body ParamSearch true "查询参数"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyList[types.SignIn] 返回签到记录信息
// @Router/signin/search [post]
func noopSignInSearch() {}

// @Summary 查询签到记录
// @Schemes
// @Description 查询签到记录
// @Tags signin
// @Param search query ParamList true "查询参数"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyList[types.SignIn] 返回签到记录信息
// @Router/signin/list [get]
func noopSignInList() {}

// @Summary 创建签到记录
// @Schemes
// @Description 创建签到记录
// @Tags signin
// @Param search body types.SignIn true "签到记录信息"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyData[types.SignIn] 返回签到记录信息
// @Router/signin/create [post]
func noopSignInCreate() {}

// @Summary 获取签到记录
// @Schemes
// @Description 获取签到记录
// @Tags signin
// @Param id path int true "签到记录ID"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyData[types.SignIn] 返回签到记录信息
// @Router/signin/{id} [get]

func noopSignInGet() {}

func signRouter(app *gin.RouterGroup) {
	app.POST("/count", curd.ApiCount[types.SignIn]())

	app.POST("/search", curd.ApiSearch[types.Recharge]())

	app.GET("/list", curd.ApiList[types.SignIn]())

	app.POST("/create", curd.ApiCreateHook[types.SignIn](nil, nil))

	app.GET("/:id", curd.ParseParamId, curd.ApiGet[types.SignIn]())
}
