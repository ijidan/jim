package pkg

import (
	"github.com/sirupsen/logrus"
	"io"
	"jim/config"
	"os"
	"sync"
)

func GetLoggerInstance(config *config.Config, root string) *logrus.Logger {
	var once sync.Once
	var instance *logrus.Logger
	once.Do(func() {
		logfile := root + "/" + config.Http.Log
		writer1 := os.Stdout
		writer2, err := os.OpenFile(logfile, os.O_WRONLY|os.O_CREATE, 0755)
		if err != nil {
			panic(err)
		}
		instance = logrus.New()
		instance.SetFormatter(&logrus.JSONFormatter{})
		instance.SetOutput(io.MultiWriter(writer1, writer2))
	})
	return instance
}
