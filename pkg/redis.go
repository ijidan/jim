package pkg

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"jim/config"
	"sync"
)

func GetRdInstance(config *config.Config) redis.Conn {
	var once sync.Once
	var instance redis.Conn
	var err error
	once.Do(func() {
		addr := fmt.Sprintf("%s:%d", config.Redis.Host, config.Redis.Port)
		instance, err = redis.Dial("tcp", addr)
		if err != nil {
			panic(err)
		}
	})
	return instance
}
