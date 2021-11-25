package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"sync"
)

// Config 配置
type Config struct {
	App struct {
		Name string `yaml:"name"`
		Env  string `yaml:"env"`
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
		Database string `yaml:"database"`
	}
	Redis struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		db       int    `yaml:"db"`
	}
}

func GetConfigInstance(root string) *Config {
	var once sync.Once
	var instance *Config
	once.Do(func() {
		instance = &Config{}
		v := viper.New()
		v.AddConfigPath(root)
		v.SetConfigName("config")
		v.SetConfigType("yaml")
		v.WatchConfig()
		v.OnConfigChange(func(in fsnotify.Event) {
		})
		if err := v.ReadInConfig(); err != nil {
			panic(err)
		}
		if err := v.Unmarshal(instance); err != nil {
			panic(err)
		}

	})
	return instance
}
