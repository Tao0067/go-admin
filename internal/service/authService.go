package service

import "paper/internal/model"

type AuthService struct {
	AdminModel *model.Admin
}

func (s AuthService) Signup(name, password string) error {
	adminUser := model.Admin{
		Name:     name,
		Password: password,
	}

	return adminUser.CreateAdmin()
}
