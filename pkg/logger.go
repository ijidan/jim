package pkg

import (
	"github.com/sirupsen/logrus"
	"jim/config"
	"sync"
)

// NewLogger log instance
func NewLogger(conf *config.Config) *logrus.Logger {
	var once sync.Once
	var instance *logrus.Logger
	once.Do(func() {
		instance = logrus.New()
		instance.SetFormatter(&logrus.JSONFormatter{})
	})
	return instance
}
