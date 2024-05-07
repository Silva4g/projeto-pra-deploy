package main

import (
	"backend/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("rodando API")

	r := router.Gerar()

	log.Fatal(http.ListenAndServe(":8000", r))
}