package handlers

import (
	"fmt"
	"log"
	"net/http"
)

type HealthCheck struct {
	l *log.Logger
}

func NewHealthCheck(l *log.Logger) *HealthCheck {
	return &HealthCheck{l}
}

func (h *HealthCheck) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.l.Println("helthy")
	// h.l.Println("User Agent:", r.UserAgent())
	h.l.Printf("%s %s %s %d\n", r.RemoteAddr, r.Method, r.URL, http.StatusOK)
	//write to response ...
	fmt.Fprintf(w, "pong")
}

// this is a fast way to create a handler - we don't want to do that as it is messy and hard to maintain
// it will use the default server and connect to it ...
// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 	log.Println("helthy")
// 	log.Println("User Agent:", r.UserAgent())
// 	data, error := ioutil.ReadAll(r.Body)
// 	query_param := r.URL.Query()

// 	if query_param.Has("a") {
// 		fmt.Printf("Has value of a: %s\n", query_param.Get("a"))
// 	}
// 	for key, val := range query_param {
// 		fmt.Printf("Query param: %s has value of: %s\n", key, val)
// 	}
// 	if error != nil {
// 		//w.WriteHeader(http.StatusBadRequest)
// 		//w.Write([]byte("Error"))
// 		// or better instead of the two lines above we do the following error :
// 		http.Error(w, "Oooops Error", http.StatusBadRequest)
// 		return
// 	}
// 	log.Printf("Data is %s", data)
// 	fmt.Fprintf(w, "Data is %s", data)

// })
