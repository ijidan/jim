package pkg

import (
	"github.com/sirupsen/logrus"
	"sync"
)

// NewLogger log instance
func NewLogger() *logrus.Logger {
	var once sync.Once
	var instance *logrus.Logger
	once.Do(func() {
		instance = logrus.New()
		instance.SetFormatter(&logrus.JSONFormatter{})
	})
	return instance
}
