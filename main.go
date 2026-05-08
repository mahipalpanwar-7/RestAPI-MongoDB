package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mahipalpanwar-7/RestAPI-MongoDB/router"
)

func main() {
	fmt.Println("MongoDB API")
	r := router.Router()
	fmt.Println("Server is getting started...")
	log.Fatal(http.ListenAndServe(":8000", r))
	fmt.Println("listening at port 8000.....")
}
