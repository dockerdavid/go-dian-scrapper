package muiscaPorts

import (
	muiscaDomain "github.com/dockerdavid/go-dian-scrapper/pkg/muisca/domain"
)

type Service interface {
	GetContributorByDocument(document string) (result *muiscaDomain.Result, err error)
}
