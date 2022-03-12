package main

import (
	"log"
	"net/http"
	"fmt"
)

const port string = ":4000"

func home(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprint(rw, "Hello from home!")
}


func main() {
	http.HandleFunc("/", home)
	fmt.Printf("Listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
