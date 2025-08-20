/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/
package cmd




import (
	"context"

	"fmt"
	"github.com/spf13/cobra"
	"time"

	"github.com/redis/go-redis/v9"
)
var userKey string 
var capacity int64
var fillRate int64

type TokenBucket struct {
	Capacity      int64
	Fillrate      int64
	CurrentTokens int64
	LastFilled    int64 // Unix timestamp
}

func TokenBucketSetUp(client *redis.Client, luaScript string, capacity int64, fillRate int64, currentTokens int64, userKey string) (bool, error) {

	ctx := context.Background()
    exists,err:=client.Exists(ctx,userKey).Result()
	if err!=nil{
		fmt.Println("Problem while checking existence")


	}
	if (exists==0){
		err := client.HSet(ctx, userKey, map[string]interface{}{
			"capacity":       capacity,
			"fillrate":       fillRate,
			"current_tokens": currentTokens,
			"last_filled":    time.Now().Unix(),
		}).Err()
		if err != nil {
			fmt.Errorf("The error while creating hashset %w", err)
	
		}

	}
    


	
	return TokenBucketExecute(client, luaScript, userKey, capacity, fillRate)
}

func TokenBucketExecute(client *redis.Client, luaScript, userKey string, capacity int64, fillRate int64) (bool, error) {
	ctx := context.Background()

	val, err := client.Eval(ctx, luaScript, []string{userKey}, []interface{}{
		capacity,
		fillRate,
		time.Now().Unix(),
		50,
	}).Result()
	if err != nil {
		panic(err)

	}

	if val.(int64) == 1 {
		return true, nil
	}else {
		return false, nil

	}
}


// tokenBucketCmd represents the tokenBucket command
var tokenBucketCmd = &cobra.Command{
	Use:   "tokenBucket",
	Short: "Uses token bucket algorithm of rate limiting",
	Long: `Uses token bucket algorithm of rate limiting  Helps with dynamic values according to the needs`,
	Run: func(cmd *cobra.Command, args []string) {
		
	},
}

func init() {
	rootCmd.AddCommand(tokenBucketCmd)

	tokenBucketCmd.Flags().String("userkey",userKey,"Sets the key set for the user for personalised storage")
	tokenBucketCmd.Flags().Int64("capacity",capacity,"Sets Capacity of bucket containing tokens to a given amount")
	tokenBucketCmd.Flags().Int64("")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// tokenBucketCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// tokenBucketCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
