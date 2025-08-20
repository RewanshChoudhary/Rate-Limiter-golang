package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/RewanshChoudhary/Rate-Limiter-golang/algorithms"
	"github.com/redis/go-redis/v9"
)
var capacity int64 
var refillRate int64
var userKey string
var currentTokens int64


var ctx = context.Background()

func main() {
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

	mux.HandleFunc("/hello",)


	scriptBytes, err := os.ReadFile("TokenBucketScript.lua")

	script:=string(scriptBytes)

	val,err:=algorithms.TokenBucketSetUp(rd,script,capacity,refillRate,currentTokens,userKey)
	

	
    if val{
		fmt.Print("Request was accepted ")


	}else {
		fmt.Println("The request is not sent")

	}
	

}
