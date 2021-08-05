package main

import (
	"flag"
	"strconv"

	"kgin/db"
	"kgin/pkg"
	"kgin/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	httpport := flag.Int("p", 8014, "HttpPort")
	flag.Parse()
	// 全局设置环境，此为开发环境，线上环境为gin.ReleaseMode
	gin.SetMode(gin.DebugMode)
	// 初始化...
	pkg.Init()
	db.Init()
	defer db.Close()
	//grpc 服务端开启
	//server.Init()

	// Initialize the routes
	router := routes.Init()
	// listen and serve on 0.0.0.0:8010
	router.Run(":" + strconv.Itoa(*httpport))
}
