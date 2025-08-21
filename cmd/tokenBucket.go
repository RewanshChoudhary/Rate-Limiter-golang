package cmd

import (
	"context"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"

	"fmt"
	"time"

	"github.com/spf13/cobra"

	"github.com/redis/go-redis/v9"
)
var userKey string 
var capacity int64
var refillRate int64
var endpoint string 
var serverport ="8080"
var testServer bool
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
		1,
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
		refillRate,err:=cmd.Flags().GetInt64("refillRate")
		if (err!=nil){
			fmt.Errorf("The following error occurred while reading the script %w",err)

		}
		
		if (err!=nil){
			fmt.Errorf("The following error occurred while getting the output fromr the function %w",err)

		}
		ep,err:=cmd.Flags().GetString("endpoint")
		create,err:=cmd.Flags().GetBool("testserver")
		
		mux:=http.NewServeMux()
		if create{


		
		port,err:=cmd.Flags().GetString("serverport")

		if (err!=nil){
			fmt.Errorf("Flag was not read hence the error %w",err)

		}
		mux.HandleFunc(ep,http.HandlerFunc(func (w http.ResponseWriter,r * http.Request){
			acc,err:=TokenBucketSetUp(rd,string(luaScript),capacity,refillRate,userKey)
			if err!=nil{
				fmt.Errorf("The error throw %w",err)

			
			}

			
			if acc {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte("Request allowed\n"))
			} else {
				w.WriteHeader(http.StatusTooManyRequests)
				w.Write([]byte("Rate limit exceeded\n"))
			}
			
		}))
		fmt.Printf("Server running on port %s, endpoint: %s\n", port, ep)

		http.ListenAndServe(":"+port,mux)
		fmt.Printf("Server running on port %s, endpoint: %s\n", port, ep)
	}else {
		resp,err:=http.Get(ep)
		if(err!=nil){
			fmt.Errorf("The error we received %w",err)

		}
		if (resp.StatusCode!=http.StatusOK){
			fmt.Println("Not a valid endpoint ")
			return
		}
		
proxy := httputil.NewSingleHostReverseProxy(&url.URL{Scheme: "https", Host: "external-service.com"})

mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    allowed, _ := TokenBucketSetUp(rd, string(luaScript), capacity, refillRate, userKey)
    if !allowed {
        w.WriteHeader(http.StatusTooManyRequests)
        w.Write([]byte("Rate limit exceeded"))
        return
    }
    proxy.ServeHTTP(w, r) // Forward request if allowed
})


	}
		
	},
}

func init() {
	rootCmd.AddCommand(tokenBucketCmd)

	tokenBucketCmd.Flags().String("userkey",userKey,"Sets the key set for the user for personalised storage")
	tokenBucketCmd.Flags().Int64("capacity",capacity,"Sets Capacity of bucket containing tokens to a given amount")
	tokenBucketCmd.Flags().Int64("refillrate",refillRate,"Sets the fillrate for your bucket configuration ")
    tokenBucketCmd.Flags().String("endpoint",endpoint,"The endpoint you want to secure with this service ")
    tokenBucketCmd.Flags().String("serverport",serverport,"Server port for testing how this works ")
tokenBucketCmd.Flags().Bool("testserver",testServer,"Provides a test server in uyour local machine to check how does the working goes")}