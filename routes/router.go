package routes

import (
	"kgin/kapi/controller"

	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	router := gin.Default()
	router.GET("/", controller.Default)
	//pprof收集cpu性能消耗
	//pprof.Register(router)
	//从DB加载策略
	//a := xormadapter.NewAdapter("mysql", "root:123456@tcp(127.0.0.1:3306)/k12_studio_v2?charset=utf8", true)
	//e := casbin.NewEnforcer("./conf/rbac_models.conf", a)
	//e.LoadPolicy()
	//new(controller.Groups).Groups(router, e)
	//router.Use(middleware.Authorize(e))
	api_v1 := router.Group("/")
	{
		new(controller.UserController).User(api_v1)
	}
	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, "找不到路由呀!大兄弟!")
	})
	return router
}
