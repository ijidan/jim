package main

import (
	"fmt"
	"github.com/spf13/cast"
	"jim/internal/global"
	"jim/internal/http/router"
	"os"
)

func main() {
	defer global.Close()
	http := global.Config.Http
	r := router.NewGin(global.Config)
	addr := fmt.Sprintf("%s:%d", http.Host, http.Port)

	fmt.Println("|-----------------------------------|")
	fmt.Println("|                jim                |")
	fmt.Println("|-----------------------------------|")
	fmt.Println("|     Http Server Start Successful  |")
	fmt.Println("|     Port:  " + cast.ToString(http.Port) + "        			|")
	fmt.Println("|     Pid:   " + fmt.Sprintf("%d", os.Getpid()) + "        			|")
	fmt.Println("|-----------------------------------|")
	fmt.Println("")

	_ = r.Run(addr)
}
