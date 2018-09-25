// 最简单的echo服务，计数器

package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)


var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler ...
func handler(w http.ResponseWriter, r *http.Request)  {
	mu.Lock()
	count ++
	mu.Unlock()
	fmt.Fprintf(w,"URL.Path = %q\n", r.URL.Path)
	fmt.Printf("%d\n", count)	// 访问/count时，这里也会执行

}

func counter(w http.ResponseWriter, r *http.Request)  {
	mu.Lock()
	fmt.Fprintf(w,"Count %d\n", count)
	mu.Unlock()
}
