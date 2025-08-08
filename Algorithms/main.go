package algorithms

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func requestHandler(tb *Tokenbucket, h http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if tb.Allow(1) {

			fmt.Println("Requests accepted")
			h.ServeHTTP(w, r)

		} else {
			fmt.Errorf("The request was failed")

		}

	})

}

func main() {
	var wg sync.WaitGroup
	defer wg.Done()

	bucket := Tokenbucket{
		Capacity:      100,
		Fillrate:      5,
		CurrentTokens: 100,
		LastFilled:    time.Now(),
	}

	mux := http.NewServeMux()

	go http.ListenAndServe(":8081", mux)

	for i:=0;i<25;i++{
		go func (id int){
			
		}
	}

}
