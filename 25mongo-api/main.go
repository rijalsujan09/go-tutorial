package main

import (
	"fmt"
	"net/http"

	"github.com/rijalsujan09/mongo-api/router"
)

func main() {
	fmt.Println("Mongo DB API")

	r := router.Router()
	fmt.Println("server  is getting started...")
	http.ListenAndServe(":4000", r)
	fmt.Println("Listening at port : 4000")
}
