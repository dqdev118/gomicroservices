package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Hello World")

		d, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(rw, "Oops", http.StatusBadRequest)

			return
		}
		// log.Printf("Data %s\n", d)
		fmt.Fprintf(rw, "Hello %s", d)
	})

	http.ListenAndServe(":8080", nil)
}
