package cmd

import (
	"fmt"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"jim/global"
	"jim/internal/http/router"
	"os"
)

// startHttpCmd represents the startHttp command
var startHttpCmd = &cobra.Command{
	Use:   "startHttp",
	Short: "start http server",
	Long:  `start http server`,
	Run: func(cmd *cobra.Command, args []string) {
		defer func() {
			pkg.Close()
		}()
		http := pkg.Conf.Http
		r := router.NewGin(pkg.Conf)
		addr := fmt.Sprintf("%s:%d", http.Host, http.Port)

		fmt.Println("|-----------------------------------|")
		fmt.Println("|                jim                |")
		fmt.Println("|-----------------------------------|")
		fmt.Println("|     Http Server Start Successful  |")
		fmt.Println("|     Port:  " + cast.ToString(http.Port) + "        			|")
		fmt.Println("|     Pid:   " + fmt.Sprintf("%d", os.Getpid()) + "        			|")
		fmt.Println("|-----------------------------------|")
		fmt.Println("")
		if err := r.Run(addr); err != nil {
			println(err.Error())
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(startHttpCmd)
}
