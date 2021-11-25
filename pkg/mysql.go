package pkg

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"jim/config"
	"sync"
)

var (
	onceDb     sync.Once
	instanceDb *gorm.DB
)

func GetDbInstance(conf *config.Config) *gorm.DB {
	onceDb.Do(func() {
		var err error
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local", conf.Mysql.User, conf.Mysql.Password, conf.Mysql.Host, conf.Mysql.Port, conf.Mysql.Database, conf.Mysql.Charset)
		instanceDb, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(err)
		}
	})
	return instanceDb
}
