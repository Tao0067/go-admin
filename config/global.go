package config

import (
	"gorm.io/gorm"
	"paper/pkg/logger"
)

var (
	ServiceConfig *Config
	DB            *gorm.DB
	Logger        *logger.Logger
)
