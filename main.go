package algorithms

import (
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

func RequestMiddlewareHandler(tb *Tokenbucket, h http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if tb.Allow(1) {

			fmt.Println("Requests accepted")
			h.ServeHTTP(w, r)

		} else {
			fmt.Errorf("The request was failed")

		}

	})

}
func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hallow")
}

func main() {
	var wg sync.WaitGroup
	

	bucket := Tokenbucket{
		Capacity:      100,
		Fillrate:      5,
		CurrentTokens: 100,
		LastFilled:    time.Now(),
	}

	mux := http.NewServeMux()
	mux.Handle("/hello", RequestMiddlewareHandler(&bucket, http.HandlerFunc(sayHello)))
	wg.Wait()

	go http.ListenAndServe(":8081", mux)
	wg.Add(25)
	for i := 0; i < 25; i++ {
		go func(id int) {
			defer wg.Done()

			resp, err := http.Get("http://localhost:8081/hello")
			if err != nil {
				panic(fmt.Errorf("The following error occured during reading th response %w", err))

			}
			defer resp.Body.Close()
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				fmt.Println("The thing occured during reading ")
			}
			fmt.Printf("[Client %d] %s (status: %d)\n", id, string(body), resp.StatusCode)

			time.Sleep(time.Second * 1)

		}(i)
	}
	wg.Wait()
}
