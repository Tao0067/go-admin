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
	CreateTime int64  `json:"create_time"`
}

type ListSearchParams struct {
	CreateTimeStart int64
	CreateTimeEnd   int64
	Name            string
	Page            int
	Size            int
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
	password = models.Encrypt(password)

	tx := config.DB.Where("name = ? and password = ?", name, password).First(&a)
	if tx.Error != nil {
		return a, tx.Error
	}

	return a, nil
}

func FindAdminByName(name string) (Admin, error) {
	a := Admin{}

	tx := config.DB.Where("name = ? ", name).First(&a)
	if tx.Error != nil {
		return a, tx.Error
	}

	return a, nil
}

func AdminUserListSearch(search ListSearchParams) (list *[]Admin, total int64, err error) {
	list = &[]Admin{}
	db := config.DB
	if search.CreateTimeStart != 0 && search.CreateTimeEnd != 0 {
		db = db.Where("`created_at` BETWEEN ? AND ?", search.CreateTimeStart, search.CreateTimeEnd)
	}
	if search.Name != "" {
		db = db.Where("`name` = ?", search.Name)
	}

	db = db.Order("create_time desc")

	err = db.Model(&Admin{}).Count(&total).Error
	if err != nil {
		return
	}
	err = db.Select("id, name, create_time").Limit(search.Size).Offset((search.Page - 1) * search.Size).Find(list).Error
	return
}

func (a *Admin) EditAdmin() error {
	if a.Password != "" {
		a.Password = models.Encrypt(a.Password)
	}
	err := config.DB.Updates(a).Error
	return err
}

func (a *Admin) DelAdmin() error {
	err := config.DB.Delete(a).Error
	return err
}
