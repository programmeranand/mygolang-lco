package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorialla/mux/router"
)

func main() {
	fmt.Println("MongoDB API")
	fmt.Println("Server is getting Started....")
	r := router.Router()
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Listening at Port 4000..")

}
