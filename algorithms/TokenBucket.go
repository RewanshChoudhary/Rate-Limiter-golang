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
