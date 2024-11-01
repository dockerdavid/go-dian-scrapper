package main

import (
	"fmt"
	"github.com/dockerdavid/go-dian-scrapper/internal/muisca/adapters/services"
	_ "github.com/joho/godotenv/autoload"
	"os"
)

func main() {
	muiscaSrv := muiscaServices.Service{}

	getContributorByDocument, err := muiscaSrv.GetContributorByDocument(os.Getenv("DOCUMENT"))

	if err != nil {
		return
	}

	fmt.Println(getContributorByDocument)
}
