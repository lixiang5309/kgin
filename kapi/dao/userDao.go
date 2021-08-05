package dao

import (
	"errors"
	"fmt"

	"kgin/db"
	"kgin/kapi/types"
	"kgin/utils"
)

type UserDao struct {
}

func (p *UserDao) Login(username, password string) error {
	sql := `SELECT id,password FROM users WHERE username=?`
	res, err := db.Db.QueryString(sql, username)
	if err != nil || len(res) == 0 {
		return errors.New("登陆失败")
	}
	if !utils.CheckPassword(res[0]["password"], password) {
		return errors.New("密码错误")
	}
	return nil
}

func (p *UserDao) Import(a []types.User) error {
	db.Db.ShowSQL(true)
	_, err := db.Db.Table("user").Insert(&a)
	fmt.Println(err)
	if err != nil {
		return errors.New("插入失败")
	}
	return nil
}

func (p *UserDao) GetUser(id int) (types.User, error) {
	a := new(types.User)
	_, err := db.Db.Table("user").Where("id=?", id).Get(&a)
	fmt.Println(err)
	if err != nil {
		return *a, errors.New("查询失败")
	}
	return *a, nil
}

func (p *UserDao) GetAllUser() ([]types.User, error) {
	a := make([]types.User, 0)
	_, err := db.Db.Table("user").Get(a)
	fmt.Println(err)
	if err != nil {
		return a, errors.New("查询失败")
	}
	return a, nil
}
