package config

import (
	"github.com/spf13/viper"
	"jim/internal/global"
)

var Config config

// Config 配置
type config struct {
	App struct {
		Name        string `yaml:"name"`
		environment string `yaml:"environment"`
	}
	Http struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
		Log  string `yaml:"log"`
	}
	Websocket struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
		Log  string `yaml:"log"`
	}
	Rpc struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
		Log  string `yaml:"log"`
	}
	Mysql struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Charset  string `yaml:"charset"`
	}
	Redis struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		Password string `yaml:"password"`
		db       int    `yaml:"db"`
	}
}

func LoadConfigFromYaml()error  {
	v:=viper.New()
	v.AddConfigPath(global.ROOT)
	v.SetConfigName("config")
	v.SetConfigType("yaml")

	if err:=v.ReadInConfig();err!=nil{
		return err
	}
	return  nil
}

//func LoadConfigFromYaml (c *config) error  {
//	v:=viper.New()
//	v.AddConfigPath()
//	v.SetConfigName("config")
//	v.set
//	v
//	c.v = viper.New();
//
//	//设置配置文件的名字
//	c.v.SetConfigName("config")
//
//	//添加配置文件所在的路径,注意在Linux环境下%GOPATH要替换为$GOPATH
//	c.v.AddConfigPath("%GOPATH/src/")
//	c.v.AddConfigPath("./")
//
//	//设置配置文件类型
//	c.v.SetConfigType("yaml");
//
//	if err := c.v.ReadInConfig(); err != nil{
//		return  err;
//	}
//
//	log.Printf("age: %s, name: %s \n", c.v.Get("information.age"), c.v.Get("information.name"));
//	return nil;
//}
