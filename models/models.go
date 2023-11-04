package models

import (
	"github.com/lwinmgmg/linux-http/env"
	"github.com/lwinmgmg/linux-http/services"
	"gorm.io/gorm"
)

var (
	Env env.Settings = env.NewEnv()
	DB  *gorm.DB     = services.DB
)

func init() {
	if err := DB.AutoMigrate(&Key{}); err != nil {
		panic(err)
	}
}
