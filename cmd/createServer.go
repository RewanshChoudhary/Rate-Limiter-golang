/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	

	"github.com/redis/go-redis/v9"
	"github.com/spf13/cobra"
)


// createServerCmd represents the createServer command
var createServerCmd = &cobra.Command{
	Use:   "createServer",
	Short: "Creates a server for testing by providing various limits",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx:=context.Background()
		rd := redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		})
	
		pong, err := rd.Ping(ctx).Result()
		if err != nil {
			panic(err)
		}
		fmt.Println(pong)
	
		
	
		
		
	},
}

func init() {
	rootCmd.AddCommand(createServerCmd)

	createServerCmd.Flags().String("serverport","9090","For a customised server port to access")

	


}
