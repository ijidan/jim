package global

import (
	"github.com/sirupsen/logrus"
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
		instance = &global{
			Root:   dir,
			Config: config.NewConfig(dir),
			Logger: pkg.NewLogger(),
		}
	})
	return instance
}
