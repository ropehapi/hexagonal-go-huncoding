package output

import (
	"github.com/ropehapi/hexagonal-hun-coding/application/domain"
	"github.com/ropehapi/hexagonal-hun-coding/configuration/rest_err"
)

type NewsPort interface {
	GetNewsPort(domain.NewsReqDomain) (*domain.NewsDomain, *rest_err.RestErr)
}