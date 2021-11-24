package global

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"jim/config"
	"jim/pkg"
	"os"
	"sync"
)

// global
type global struct {
	Root   string
	Config *config.Config
	Logger *logrus.Logger
	Db     *gorm.DB
}

// NewGlobal global instance
func NewGlobal() *global {
	var once sync.Once
	var instance *global
	once.Do(func() {
		dir, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		config := config.NewConfig(dir)
		instance = &global{
			Root:   dir,
			Config: config,
			Logger: pkg.NewLogger(config),
			Db:     pkg.NewDb(config),
		}
	})
	return instance
}
