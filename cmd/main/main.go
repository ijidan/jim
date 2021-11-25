package main

import (
	"fmt"
	"jim/internal/global"
	"jim/internal/http/router"
)

func main() {
	defer global.Close()
	http := global.Config.Http
	r := router.NewGin()
	addr := fmt.Sprintf("%s:%d", http.Host, http.Port)
	_ = r.Run(addr)
}
