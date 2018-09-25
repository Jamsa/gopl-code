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
	
	fmt.Printf("%d\n", count)	// 访问/count时，这里也会执行


	fmt.Fprintf(w,"%s %s %s\n",r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k,v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}

	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k,v)
	}
}

func counter(w http.ResponseWriter, r *http.Request)  {
	mu.Lock()
	fmt.Fprintf(w,"Count %d\n", count)
	mu.Unlock()
}
