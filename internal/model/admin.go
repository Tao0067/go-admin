package model

import (
	"paper/config"
	"paper/pkg/models"
	"time"
)

type Admin struct {
	Id         int    `json:"id"`
	Uuid       string `json:"uuid"`
	Name       string `json:"name"`
	Password   string `json:"password"`
	CreateTime int64  `json:"create_at"`
}

func (a *Admin) TableName() string {
	return "xcx_admin_user"
}

// CreateAdmin 创建用户
func (a *Admin) CreateAdmin() error {
	a.CreateTime = time.Now().Unix()
	a.Password = models.Encrypt(a.Password)
	a.Uuid = models.CreateUUID()
	tx := config.DB.Create(a)

	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

// FindAdminByNameAndPaw 查询用户
func FindAdminByNameAndPaw(name string, password string) (Admin, error) {
	a := Admin{}

	tx := config.DB.Where("name = ? amd password = ?", name, password).First(&a)
	if tx.Error != nil {
		return a, tx.Error
	}

	return a, nil
}
