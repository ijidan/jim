package main

import (
	"github.com/spf13/cast"
	"jim/internal/global"
	"jim/internal/router"
)

func main() {
	http := global.NewGlobal().Config.Http
	host := http.Host
	port := http.Port
	r := router.NewGin()
	_ = r.Run(host + ":" + cast.ToString(port))
}
