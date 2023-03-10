package main

import (
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("started api...")

	r := router.GenerateRouter()

	log.Fatal(http.ListenAndServe(":5000", r))
}
