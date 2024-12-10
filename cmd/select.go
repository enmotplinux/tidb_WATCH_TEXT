/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"tidb-go/tidb/db"
)
var(  
	ip string
	port string

)

// selectCmd represents the select command
var selectCmd = &cobra.Command{
	Use:   "select",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		db := &db.Database{}
		err := db.Init_db(ip, port )
		if err != nil {
			fmt.Println(err)
		}
		err = db.QueryRowDemo()
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(selectCmd)
	selectCmd.Flags().StringVarP(&ip, "ip", "i", "127.0.0.1", "ip")
	selectCmd.Flags().StringVarP(&port, "port", "p", "4000", "port")


}

/*
add 
/users/user/go/bin/cobra-cli add select
user@EBJ1206463 tidb-go % go run main.go select --ip "192.168.164.195"  --port 4000
select called
ip 192.168.164.195
port 4000

*/