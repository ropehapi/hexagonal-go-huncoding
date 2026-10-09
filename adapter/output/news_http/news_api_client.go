package news_http

import (
	resty "github.com/go-resty/resty/v2"
	"github.com/jinzhu/copier"
	"github.com/ropehapi/hexagonal-hun-coding/adapter/output/model/response"
	"github.com/ropehapi/hexagonal-hun-coding/application/domain"
	"github.com/ropehapi/hexagonal-hun-coding/configuration/env"
	"github.com/ropehapi/hexagonal-hun-coding/configuration/rest_err"
)

type newsClient struct {}

func NewNewsClient() *newsClient{
	client = resty.New().SetBaseURL("https://newsapi.org/v2")
	return &newsClient{}
}

var (client *resty.Client)

func (nc *newsClient) GetNewsPort(newDomain domain.NewsReqDomain) (*domain.NewsDomain, *rest_err.RestErr){	
	newsResponse := &response.NewsClientResponse{}
	
	_, err := client.R().SetQueryParams(map[string]string{
		"q": newDomain.Subject,
		"from": newDomain.From,
		"apiKey": env.GetNewsTokenApi(),
	}).SetResult(newsResponse).Get("/everything")
	
	if err != nil {
		return nil, rest_err.NewInternalServerError("Error trying to call NewsAPI with params")
	}

	newsResponseDomain := &domain.NewsDomain{}
	copier.Copy(newsResponseDomain, newsResponse)

	return newsResponseDomain, nil
}