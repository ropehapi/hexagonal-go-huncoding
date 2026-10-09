package input

import (
	"github.com/ropehapi/hexagonal-hun-coding/application/domain"
	"github.com/ropehapi/hexagonal-hun-coding/configuration/rest_err"
)

type NewsUseCase interface {
	GetNewsService(domain.NewsReqDomain) (*domain.NewsDomain, *rest_err.RestErr)
}