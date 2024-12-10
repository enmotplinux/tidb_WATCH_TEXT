/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"
	"tidb-go/tidb/db"

	"github.com/spf13/cobra"
)

var (

	pid  string
)

// UnbindCmd represents the Unbind command
var UnbindCmd = &cobra.Command{
	Use:   "Unbind",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		db := &db.Database{}
		db.Init_db(ip, port)
		// 尝试将 pid 从 string 转换为 int
		intPid, err := strconv.Atoi(pid)
		if err != nil {
			fmt.Println("Invalid pid:", err)
			return
		}
		// 调用 BindsqlUnbind 函数
		err = db.BindsqlUnbind(intPid)
		fmt.Println(err)


	},
}

func init() {
	rootCmd.AddCommand(UnbindCmd)
	UnbindCmd.Flags().StringVarP(&ip, "ip", "i", "127.0.0.1", "ip")
	UnbindCmd.Flags().StringVarP(&port, "port", "p", "4000", "port")
	UnbindCmd.Flags().StringVarP(&pid, "pid", "q", "1", "pid")
}
/*

 解绑：
 go run main.go Unbind --ip "10.178.65.35" --port 7021 --pid 6  id 是查询出来的
查询：
 go run main.go select --ip "10.178.65.35" --port 7021
 拉黑：
 go run main.go block --ip "10.178.65.35" --port 7021 --sql "select * from test.t1"   
 go run main.go block --ip "10.178.65.35" --port 7021 --sql "select * from test.t1 limit 1"   


*/