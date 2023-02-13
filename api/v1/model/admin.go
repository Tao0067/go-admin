package model

import (
	"crypto/md5"
	"fmt"
	"paper/pkg/models"
	"time"
)

type admin struct {
	Id       int       `json:"id"`
	Name     string    `json:"name"`
	Password string    `json:"password"`
	CreateAt time.Time `json:"create_at"`
}

func (a *admin) CreateAdmin() error {
	a.CreateAt = time.Time{}
	sum := md5.Sum([]byte(a.Password))
	a.Password = fmt.Sprintf("%x", sum)
	tx := models.DB.Create(a)

	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func FindAdminByNameAndPaw(name string, password string) (admin, error) {
	a := admin{}

	tx := models.DB.Where("name = ? amd password = ?", "name", password).First(&a)
	if tx.Error != nil {
		return a, tx.Error
	}

	return a, nil
}
