package main

import (
	"context"
	"fmt"

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

	err=rd.Set(ctx,"Name","Blah",0).Err()

    if (err!=nil){
		panic(err);

	}
	val,er:=rd.Get(ctx,"Name").Result()
	if (er!=nil){
		panic(er)
	}
	fmt.Print(val)
	

}
