package controller

import (
	"github.com/gin-gonic/gin"
)

type ApiController struct {
}

func (p *ApiController) DataReturn(c *gin.Context, err error, data interface{}) {
	if err != nil {
		c.JSON(400, gin.H{
			"code": 10001,
			"msg":  err.Error(),
			"data": nil,
		})
	} else {
		c.JSON(200, gin.H{
			"code": 10000,
			"msg":  "success",
			"data": data,
		})
	}
}

func (p *ApiController) DataReturnErr(c *gin.Context, err error) {
	c.JSON(400, gin.H{
		"code": 10001,
		"msg":  err.Error(),
		"data": nil,
	})
}

func (p *ApiController) DataReturnParameterErr(c *gin.Context) {
	c.JSON(400, gin.H{
		"code": 10001,
		"msg":  "参数错误",
		"data": nil,
	})
}
