package main

import (
	"fmt"
	"net/http"
	"os"
	"github.com/sirupsen/logrus"
)

func main() {
	var log = logrus.New()
	log.Out = os.Stdout
	http.HandleFunc("/helloworld", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})
	fmt.Printf("Server running (port=8080), route: http://localhost:8080/helloworld\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
