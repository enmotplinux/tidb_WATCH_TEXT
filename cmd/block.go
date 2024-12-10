/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"tidb-go/tidb/db"

	"github.com/spf13/cobra"
)

var(  

	sql string

)

// blockCmd represents the block command
var blockCmd = &cobra.Command{
	Use:   "block",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		db := &db.Database{}
		db.Init_db(ip, port)
		err := db.Bindsql(sql)
		fmt.Println(err)
		
	},
}

func init() {
	rootCmd.AddCommand(blockCmd)
	blockCmd.Flags().StringVarP(&ip, "ip", "i", "127.0.0.1", "ip")
	blockCmd.Flags().StringVarP(&port, "port", "p", "4000", "port")
	blockCmd.Flags().StringVarP(&sql, "sql", "s", "select", "sql")


	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// blockCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// blockCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

/*
go run main.go block --ip "127.0.0.1" --port 4000   --sql "select * from test.t1 "       
<nil>
*/