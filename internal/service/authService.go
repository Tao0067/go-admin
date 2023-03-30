package service

import (
	"errors"
	"paper/config"
	"paper/internal/model"
)

type AuthService struct {
	AdminModel *model.Admin
}

func (s AuthService) Signup(name, password string) error {
	admin, _ := model.FindAdminByName(name)

	if admin.Id > 0 {
		return errors.New("用户名 " + name + " 已经存在！")
	}

	adminUser := model.Admin{
		Name:     name,
		Password: password,
	}

	return adminUser.CreateAdmin()
}

func (a AuthService) Login(name, password string) (model.Admin, error) {
	adminUser, err := model.FindAdminByNameAndPaw(name, password)
	if err != nil || adminUser.Id == 0 {
		config.Logger.Error("用户不存在", name, err.Error())
		return adminUser, errors.New("用户名 " + name + " 已经存在！")
	}

	return adminUser, nil
}
