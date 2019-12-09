package userInterface

import (
	"bytes"
	"github.com/wuxiaoxiaoshen/stackoverflowCli/app/domain/model"
	"github.com/wuxiaoxiaoshen/stackoverflowCli/app/infrastructure"
	"github.com/wuxiaoxiaoshen/stackoverflowCli/app/infrastructure/errors"

	"github.com/PuerkitoBio/goquery"
)

type Reputation struct {
	model.Reputation
}

type ReputationResult []Reputation

func (r *Reputation) FindAll(url string) (interface{}, error) {
	// 1. one step : get response
	response, err := infrastructure.ResponseStorager(url)
	if err != nil {
		return nil, &errors.ErrorResponse
	}

	var reputationResult ReputationResult
	// 2. two step : get content
	doc, _ := goquery.NewDocumentFromReader(bytes.NewReader(response))
	doc.Find("#user-browser > div > div").Each(func(i int, selection *goquery.Selection) {
		userName := selection.Find(".user-details").Find("a").Text()
		userLocation := selection.Find(".user-details").Find("span.user-location").Text()
		var userFlair string
		selection.Find(".user-details").Find(".-flair > span").Each(func(i int, selection *goquery.Selection) {
			temp, _ := selection.Attr("title")
			userFlair += ", " + temp
		})

		//fmt.Println(userName, userLocation, userFlair)
		if i <= 10 && len(reputationResult) <= 10 {
			oneReputation := Reputation{
				*model.NewReputation(userName, []string{}, userLocation, userFlair),
			}
			reputationResult = append(reputationResult, oneReputation)
		}

	})
	infrastructure.Marshal(reputationResult)
	return nil, nil
}

func (r *Reputation) Single(url string, number int) (interface{}, error) {
	return Reputation{}, nil
}
