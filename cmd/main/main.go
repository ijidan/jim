package main

import (
	"fmt"
	"jim/internal/global"
	"jim/internal/router"
)

func main() {
	http := global.NewGlobal().Config.Http
	r := router.NewGin()
	addr := fmt.Sprintf("%s:%d", http.Host, http.Port)
	_ = r.Run(addr)
}
