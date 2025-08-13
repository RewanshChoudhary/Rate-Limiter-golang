package algorithms

import (
	"context"

	"fmt"

	"time"

	"github.com/redis/go-redis/v9"
)


type TokenBucket struct {
    Capacity      int64
    Fillrate      int64
    CurrentTokens int64
    LastFilled    int64 // Unix timestamp
}




func TokenBucketSetUp(client *redis.Client,luaScript string )(string,error){
	ctx:=context.Background()
	userKey:="User12345"



	err:=client.HSet(ctx ,userKey,map[string]interface{}{
		"capacity":       100,
		"fillrate":       10,
		"current_tokens": 80,
		"last_filled":    time.Now().Unix(),

	}).Err()

	if (err!=nil){
		fmt.Errorf("The error while creating hashset %w",err)

	}


	




}