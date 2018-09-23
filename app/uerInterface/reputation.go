package uerInterface

import (
	"bytes"
	"fmt"
	"stackOverFlowClient/app/domain/model"
	"stackOverFlowClient/app/infrastructure"
	"stackOverFlowClient/app/infrastructure/errors"

	"github.com/PuerkitoBio/goquery"
)

type Reputation struct {
	model.Reputation
}

type Reputations []Reputation

func (r *Reputation) FindAll(url string) (interface{}, error) {
	// 1. one step : get response
	response, err := infrastructure.ResponseStorager(url)
	if err != nil {
		return nil, &errors.ErrorResponse
	}

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

		fmt.Println(userName, userLocation, userFlair)

	})

	return Reputations{}, nil
}

func (r *Reputation) Single(url string, number int) (interface{}, error) {
	return Reputation{}, nil
}
