package muiscaPorts

import (
	muiscaDomain "go-dian-scrapper/internal/muisca/domain"
)

type Service interface {
	GetContributorByDocument(document string) (result *muiscaDomain.Result, err error)
}
