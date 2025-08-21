
package cmd

import (
	"context"
	"os"

	"fmt"
	"time"

	"github.com/spf13/cobra"

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

func TokenBucketSetUp(client *redis.Client, luaScript string, capacity int64, fillRate int64, userKey string) (bool, error) {

	ctx := context.Background()
    exists,err:=client.Exists(ctx,userKey).Result()
	if err!=nil{
		fmt.Println("Problem while checking existence")


	}
	if (exists==0){
		err := client.HSet(ctx, userKey, map[string]interface{}{
			"capacity":       capacity,
			"fillrate":       fillRate,
			"current_tokens": capacity,
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
func returnStatusCode(acc bool ){

	// Will change for future additions

if (acc){
	fmt.Println("Request sent")

}else {
	fmt.Println("Denied")


}
}
// tokenBucketCmd represents the tokenBucket command
var tokenBucketCmd = &cobra.Command{
	Use:   "tokenBucket",
	Short: "Uses token bucket algorithm of rate limiting",
	Long: `Uses token bucket algorithm of rate limiting  Helps with dynamic values according to the needs`,
	Run: func(cmd *cobra.Command, args []string) {
		
		rd := redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		})
		luaScript,err:=os.ReadFile("TokenBucketScript.lua")
		if (err!=nil){
			fmt.Errorf("The following error occurred while reading the script %w",err)

		}
		userKey,err:=cmd.Flags().GetString("userkey")
		if (err!=nil){
			fmt.Errorf("The following error occurred while reading the script %w",err)

		}
		capacity,err:=cmd.Flags().GetInt64("capacity")
		if (err!=nil){
			fmt.Errorf("The following error occurred while reading the script %w",err)

		}
		refillRate,err:=cmd.Flags().GetInt64("fillRate")
		if (err!=nil){
			fmt.Errorf("The following error occurred while reading the script %w",err)

		}
		accepted,err:=TokenBucketSetUp(rd,string(luaScript),capacity,refillRate,userKey)
		if (err!=nil){
			fmt.Errorf("The following error occurred while getting the output fromr the function %w",err)

		}

		returnStatusCode(accepted)

		
	},
}

func init() {
	rootCmd.AddCommand(tokenBucketCmd)

	tokenBucketCmd.Flags().String("userkey",userKey,"Sets the key set for the user for personalised storage")
	tokenBucketCmd.Flags().Int64("capacity",capacity,"Sets Capacity of bucket containing tokens to a given amount")
	tokenBucketCmd.Flags().Int64("refillrate",fillRate,"Sets the fillrate for your bucket configuration ")


}
