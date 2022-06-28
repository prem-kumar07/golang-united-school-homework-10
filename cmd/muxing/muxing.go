package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

/**
Please note Start functions is a placeholder for you to start your own solution.
Feel free to drop gorilla.mux if you want and use any other solution available.

main function reads host/port from env just for an example, flavor it following your taste
*/

// Start /** Starts the web server listener on given host and port.
func Start(host string, port int) {
	router := mux.NewRouter()

	router.HandleFunc("/",func(rw http.ResponseWriter, r *http.Request) {
		rw.WriteHeader(200)
	}).Methods("GET")

	router.HandleFunc("/name/{PARAM}",func(rw http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
        param := vars["PARAM"]
        fmt.Fprintf(rw,"Hello, %s!",param)
	}).Methods("GET")

	router.HandleFunc("/bad",func(rw http.ResponseWriter, r *http.Request) {
		rw.WriteHeader(500)
	})
	router.HandleFunc(`/data`   ,func(rw http.ResponseWriter, r *http.Request) {
		data,err:=ioutil.ReadAll(r.Body)
		if err !=nil {
			rw.WriteHeader(500)
		}
		fmt.Fprintf(rw,"I got message:\n%s",data)
		
	}).Methods("POST")
	router.HandleFunc(`/headers`   ,func(rw http.ResponseWriter, r *http.Request) {
	a, err := strconv.Atoi(r.Header.Get("a"))
	if err != nil {
		rw.WriteHeader(500)
		return
	}
	b, err := strconv.Atoi(r.Header.Get("b"))
	if err != nil {
		rw.WriteHeader(500)
		return
	}
	val:=a+b
	rw.Header().Set("a+b",fmt.Sprint(val))

	}).Methods("POST")
   


	log.Println(fmt.Printf("Starting API server on %s:%d\n", host, port))
	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), router); err != nil {
		log.Fatal(err)
	}
}

//main /** starts program, gets HOST:PORT param and calls Start func.
func main() {
	host := os.Getenv("HOST")
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 8081
	}
	Start(host, port)
}
