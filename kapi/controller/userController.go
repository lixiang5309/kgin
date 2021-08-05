package controller

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"kgin/kapi/service"
	"kgin/kapi/types"
	"kgin/tools"
	"kgin/utils"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	ApiController
}

func (p *UserController) User(parentRoute *gin.RouterGroup) {
	router := parentRoute.Group("/users")
	router.GET("login", p.Login)
	router.POST("set_support", p.SetSupport)
	router.GET("get_support", p.GetSupport)
	//router.Use(middleware.Auth())
	router.GET("import", p.Import)

}

var userService = new(service.UserService)

func (p *UserController) Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	err := userService.Login(username, password)
	p.DataReturn(c, err, nil)
}

func (p *UserController) Import(c *gin.Context) {
	file, err := os.Open("./file/test.csv")
	if err != nil {
		p.DataReturnErr(c, errors.New("打开文件失败"))
		return
	}
	defer file.Close()
	content := make([]types.User, 0)
	fileContent := make([]string, 0, 3)
	reader := csv.NewReader(file)
	num := 0
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error:", err)
			return
		}
		num++
		if num == 1 {
			continue
		}
		fileContent = strings.Split(record[0], " ")
		content = append(content, types.User{Name: fileContent[0], Sex: fileContent[2], Age: fileContent[1]})
	}
	err = userService.Import(content)
	p.DataReturn(c, err, nil)
}

func (p *UserController) SetSupport(c *gin.Context) {
	user_id := c.PostForm("user_id")
	su, err := tools.SupportNew("127.0.0.1:6379", "123456", "lx_34215")
	if err != nil {
		p.DataReturnErr(c, err)
		return
	}
	defer su.Close()
	err = su.SetSupport(user_id)
	p.DataReturn(c, err, nil)
}

func (p *UserController) GetSupport(c *gin.Context) {
	user_id := c.Query("user_id")
	su, err := tools.SupportNew("127.0.0.1:6379", "123456", "lx_34215")
	if err != nil {
		p.DataReturnErr(c, errors.New("连接数据库错误"))
		return
	}
	defer su.Close()
	mapSupport := make(map[string]string, 2)
	mapSupport["total"] = utils.Int64ToStr(su.GetSupportTotal())
	mapSupport["if_support"] = "0"
	a, err := su.GetUserSupport(user_id)
	if err != nil {
		p.DataReturnErr(c, err)
		return
	}
	if a {
		mapSupport["if_support"] = "1"
	}
	p.DataReturn(c, nil, mapSupport)
}
