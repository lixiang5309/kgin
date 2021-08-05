package middleware

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"kgin/db"

	redlock_go "github.com/tengzbiao/redlock-go"

	"gopkg.in/mgo.v2/bson"

	"github.com/casbin/casbin"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

const AuthKey = "lixiang"
const Del_key = "GEAgX62RT5wcbdP7zdyr7n3lcrtVdqGSpkRiC08qQ0W"

type MyCustomClaims struct {
	Application string `json:"application"`
	UserId      int    `json:"user_id"`
	jwt.StandardClaims
}

type Claims_Del struct {
	Key string `json:"key"`
	jwt.StandardClaims
}

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("this is Auth")
		c.Next()
	}
}

func Authorize(e *casbin.Enforcer) gin.HandlerFunc {
	return func(c *gin.Context) {
		//获取请求的URL
		obj := c.Request.URL.Path
		//获取请求方法
		act := c.Request.Method
		//获取用户的角色
		sub := "admin,bb,cc"
		//判断策略中是否存在
		arr := strings.Split(sub, ",")
		casbin_flag := false
		for _, val := range arr {
			if ok := e.Enforce(val, obj, act); ok {
				casbin_flag = true
				break
			}
		}
		if casbin_flag {
			//验证通过
			c.Next()
		} else {
			c.JSON(200, gin.H{
				"status": 400,
				"data":   "",
				"msg":    "很遗憾,权限验证没有通过!",
			})
			c.Abort()
		}
	}
}

func SaveLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.Request.Header
		method := c.Request.Method
		//如果需要存返回的数据 在数据返回那边 c.Set("data",return_data)
		data, _ := c.Get("data")
		url := c.Request.URL.Path
		mdata := &db.ApiLog{
			Id:      bson.NewObjectId(),
			Url:     url,
			Project: "test",
			Header:  header,
			Method:  method,
			Date:    time.Now().Format("2006-01-02 15:04:05"),
			Data:    data,
		}
		switch method {
		case http.MethodGet:
			mdata.Param = c.Request.URL.RawQuery
		}
		err := db.Mgdb.Insert(mdata)
		if err != nil {
			log.Fatal(err)
		}
		c.Next()
	}
}

//redis 分布式锁
func RedisLocked(lockKey string, doFunc func()) {
	servers := []string{"redis://123456@127.0.0.1:6379/0"}
	lock := redlock_go.NewRedLock(servers)
	// 加锁200毫秒
	ret := lock.Lock(lockKey, 200)
	// 释放锁
	defer lock.Unlock(ret)
	doFunc()
}
