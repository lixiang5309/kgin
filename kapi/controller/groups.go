package controller

import (
	"fmt"

	"github.com/casbin/casbin"

	"github.com/gin-gonic/gin"
)

type Groups struct {
}

func (p *Groups) Groups(router *gin.Engine, e *casbin.Enforcer) {
	router.POST("group/add", func(c *gin.Context) {
		fmt.Println("增加Policy")
		if ok := e.AddPolicy(c.PostForm("user"), c.PostForm("url"), c.PostForm("method")); !ok {
			fmt.Println("Policy已经存在")
		} else {
			fmt.Println("增加成功")
		}
	})
	//删除policy
	router.POST("/group/delete", func(c *gin.Context) {
		fmt.Println("删除Policy")
		if ok := e.RemovePolicy(c.PostForm("user"), c.PostForm("url"), c.PostForm("method")); !ok {
			fmt.Println("Policy不存在")
		} else {
			fmt.Println("删除成功")
		}
	})
	//获取policy
	router.GET("/group/get", func(c *gin.Context) {
		fmt.Println("查看policy")
		list := e.GetPolicy()
		for _, vlist := range list {
			for _, v := range vlist {
				fmt.Printf("value: %s, ", v)
			}
		}
	})
}
