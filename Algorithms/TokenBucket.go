package algorithms

import (
	"fmt"
	"math"
	"sync"
	"time"
)

//capacity
//rate
//mutex
//requestConclusion
//lastFilled
//currentTokens
type Tokenbucket struct{
	Capacity float64
	Fillrate float64
	CurrentTokens float64
    LastFilled time.Time
	mu sync.Mutex

}
//token bucket algorithm 

func (bucket *Tokenbucket)Allow(n float64) bool{
	bucket.mu.Lock()

	defer bucket.mu.Unlock()
    
	elapsedTime:=time.Since(bucket.LastFilled).Seconds()
	bucket.LastFilled=time.Now()


	newTokens:=bucket.CurrentTokens+elapsedTime*bucket.Fillrate;

	bucket.CurrentTokens=math.Min(newTokens,bucket.Capacity)



	if (n>bucket.CurrentTokens){
		fmt.Println("The request was denied")
		return false


	}
	fmt.Println("Accepted")
	return true






	







}