package services

import (
	"fmt"

	"github.com/lwinmgmg/linux-http/env"
	"gorm.io/driver/sqlite" // Sqlite driver based on CGO

	// "github.com/glebarez/sqlite" // Pure go SQLite driver, checkout https://github.com/glebarez/sqlite for details
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	Env env.Settings = env.NewEnv()
)

func init() {
	var err error
	fmt.Println(Env)
	DB, err = gorm.Open(sqlite.Open(Env.LH_DB_PATH), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}
