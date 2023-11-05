package api

import (
	"github.com/gin-gonic/gin"
	"github.com/lwinmgmg/linux-http/env"
	"github.com/lwinmgmg/linux-http/middleware"
	"github.com/lwinmgmg/linux-http/models"
	"gorm.io/gorm"
)

var Env env.Settings = env.NewEnv()

type ControllerV1 struct {
	App gin.IRoutes
	DB  *gorm.DB
}

func (v1 *ControllerV1) HandleRoutes() {
	var keys []models.Key
	v1.DB.Model(models.Key{}).Find(&keys)
	keysLength := len(keys)
	keyMap := make(map[string]int, keysLength)
	for i := 0; i < keysLength; i++ {
		keyMap[keys[i].Key] = keys[i].Count
	}
	router := v1.App.Use(middleware.JwtMiddleware(keyMap))
	router.POST("/api/v1/func/linux/execute", v1.LinuxExecute)
}
