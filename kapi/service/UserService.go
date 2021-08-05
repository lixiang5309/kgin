package service

import (
	"context"
	"encoding/json"

	"kgin/kapi/dao"
	"kgin/kapi/types"
	kgrpc "kgin/kgrpc/pb"
	"kgin/utils"
)

type UserService struct {
}

var userDao = new(dao.UserDao)

func (p *UserService) Login(username, password string) error {
	return userDao.Login(username, password)
}

func (p *UserService) Import(a []types.User) error {
	return userDao.Import(a)
}

func (p *UserService) GetUser(ctx context.Context, id *kgrpc.Id) (*kgrpc.User, error) {
	list, err := userDao.GetUser(utils.StrToInt(id.Id))
	if err != nil {
		return nil, err
	}
	data, _ := json.Marshal(list)
	return &kgrpc.User{User: string(data)}, nil
}

func (p *UserService) GetAllUsers(ctx context.Context, id *kgrpc.All) (*kgrpc.Users, error) {
	list, err := userDao.GetAllUser()
	if err != nil {
		return nil, err
	}
	data, _ := json.Marshal(list)
	return &kgrpc.Users{List: string(data)}, nil
}
