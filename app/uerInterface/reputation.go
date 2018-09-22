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
	// #user-browser > div.grid-layout
	doc.Find("#user-browser > div > div").Each(func(i int, selection *goquery.Selection) {
		fmt.Println("123")
		userName := selection.Find(".user-details").Find("a").Text()
		userLocation := selection.Find(".user-details").Find("span.user-location").Text()
		userFlair, _ := selection.Find(".user-details").Find(".-flair > span").Attr("title")
		var userTags []string

		//#user-browser > div.grid-layout > div:nth-child(1) > div.user-tags > a:nth-child(1)
		fmt.Println(selection.Find(".user-tags").Text())
		selection.Find("div.user-tags > a").Each(func(i int, selection *goquery.Selection) {
			fmt.Println(selection.Text())
			fmt.Println(1234)
			userTags = append(userTags, selection.Text())
		})

		fmt.Println(userName, userLocation, userFlair, userTags)

	})

	return Reputations{}, nil
}

func (r *Reputation) Single(url string, number int) (interface{}, error) {
	return Reputation{}, nil
}
