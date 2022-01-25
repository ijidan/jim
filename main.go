package main

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cast"
	"jim/internal/http/router"
	"jim/pkg"
	"os"
)

func buildTable(host string, port uint, pid int) *tablewriter.Table {
	addr := fmt.Sprintf("%s:%d", host, port)
	data := [][]string{
		[]string{"1", "Application", "Jim"},
		[]string{"2", "Address", addr},
		[]string{"3", "Port", cast.ToString(port)},
		[]string{"4", "Pid", fmt.Sprintf("%d", pid)},
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Id", "Info", "Desc"})
	table.AppendBulk(data)
	table.SetAlignment(tablewriter.ALIGN_LEFT) // Set Alignment
	return table
}

func main() {
	defer func() {
		pkg.Close()
	}()
	http := pkg.Conf.Http
	r := router.NewGin(pkg.Conf)
	addr := fmt.Sprintf("%s:%d", http.Host, http.Port)
	// Send output
	table := buildTable(http.Host, http.Port, os.Getpid())
	table.Render()
	if err := r.Run(addr); err != nil {
		println(err.Error())
		os.Exit(1)
	}
}
