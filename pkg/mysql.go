package pkg

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"jim/config"
	"sync"
)

func NewDb(conf *config.Config) *gorm.DB {
	var once sync.Once
	var instance *gorm.DB
	once.Do(func() {
		var err error
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local", conf.Mysql.User, conf.Mysql.Password, conf.Mysql.Host, conf.Mysql.Port, conf.Mysql.Database, conf.Mysql.Charset)
		instance, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(err)
		}
	})
	return instance
}
