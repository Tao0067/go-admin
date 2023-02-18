package model

import (
	"paper/config"
	"time"
)

type Sessions struct {
	Id         uint   `json:"id"`
	Uuid       string `json:"uuid"`
	UserId     uint   `json:"user_id"`
	CreateTime int64  `json:"create_time"`
}

func (s Sessions) CreateSessions() error {
	s.CreateTime = time.Now().Unix()
	return config.DB.Create(s).Error
}

func (s Sessions) delSessions(id int) error {
	return config.DB.Where("id = ?", id).Delete(s).Error
}

func FindSessions(id int) (*Sessions, error) {
	var s Sessions
	tx := config.DB.Where("id = ?", id).First(&s)
	return &s, tx.Error
}
