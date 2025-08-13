package main

import (
	"context"
	"fmt"
	"os"



	"github.com/redis/go-redis/v9"
)

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


	scriptBytes, err := os.ReadFile("TokenBucketScript.lua")

	script:=string(scriptBytes)

	





	

}
