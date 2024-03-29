package main

import (
	"api/cmd/config"
	"api/cmd/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("started api...")
	config.Config()

	r := router.GenerateRouter()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.PORT), r))
}
