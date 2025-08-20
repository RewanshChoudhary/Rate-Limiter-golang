/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"net/http"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/cobra"
)
var serverport string 

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
	
		
	
		mux:=http.NewServeMux()
		port,err:=cmd.Flags().GetString("serverport")

		if (err!=nil){
			fmt.Errorf("Flag was not read hence the error %w",err)

		}

	
		
		
	
		http.ListenAndServe(":"+port,mux)

	
		
	},
}

func init() {
	rootCmd.AddCommand(createServerCmd)

	createServerCmd.Flags().String("serverport",serverport,"For a customised server port to access")

	

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createServerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createServerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
