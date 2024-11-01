package main

import (
	"fmt"
	_ "github.com/joho/godotenv/autoload"
	"go-dian-scrapper/internal/muisca/adapters/services"
	"os"
)

func main() {
	muiscaSrv := services.Service{}

	getContributorByDocument, err := muiscaSrv.GetContributorByDocument(os.Getenv("DOCUMENT"))

	if err != nil {
		return
	}

	fmt.Println(getContributorByDocument)
}
