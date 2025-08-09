package main

import (
	"fmt"
	"github.com/RewanshChoudhary/Rate-Limiter-golang/algorithms"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"sync"
	"time"
)

func RequestMiddlewareHandler(tb *algorithms.TokenBucket, h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		res, err := tb.Allow(1)
		if err != nil {
			http.Error(w, err.Error(), http.StatusTooManyRequests)
			return
		}
		if res {
			fmt.Println("Request accepted")
			h.ServeHTTP(w, r)
		} else {
			http.Error(w, "Request denied", http.StatusBadRequest)
		}
	})
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello")
	w.Write([]byte("Hello from server"))
}

func test(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	count := vars["workers"]
	fmt.Println("Workers count:", count)
	w.Write([]byte(fmt.Sprintf("Workers: %s", count)))
}

func main() {
	var wg sync.WaitGroup

	bucket := algorithms.TokenBucket{
		Capacity:      100,
		Fillrate:      5,
		CurrentTokens: 100,
		LastFilled:    time.Now(),
	}

	
	r := mux.NewRouter()

	
	r.Handle("/hello", RequestMiddlewareHandler(&bucket, http.HandlerFunc(sayHello)))
	r.Handle("/test/{workers}", RequestMiddlewareHandler(&bucket, http.HandlerFunc(test)))

	go http.ListenAndServe(":8081", r)

	wg.Add(200)
	for i := 0; i < 200; i++ {
		go func(id int) {
			defer wg.Done()

			resp, err := http.Get("http://localhost:8081/hello")
			if err != nil {
				panic(fmt.Errorf("error making request: %w", err))
			}
			defer resp.Body.Close()

			body, _ := io.ReadAll(resp.Body)
			fmt.Printf("[Client %d] %s (status: %d)\n", id, string(body), resp.StatusCode)

			time.Sleep(time.Second * 10)
		}(i)
	}
	wg.Wait()
}
