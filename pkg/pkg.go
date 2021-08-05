package pkg

import (
	"kgin/pkg/redis"
)

//初始化服务器组件和sdk
func Init() {
	redis.Newredis()
}
