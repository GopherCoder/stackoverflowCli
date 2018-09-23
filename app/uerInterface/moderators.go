package uerInterface

import (
	"bytes"
	"fmt"
	"stackOverFlowClient/app/domain/model"
	"stackOverFlowClient/app/infrastructure"

	"github.com/PuerkitoBio/goquery"
)

type Moderators struct {
	model.Moderators
}

func (m *Moderators) FindAll(url string) (interface{}, error) {
	response, _ := infrastructure.ResponseStorager(url)
	fmt.Println(string([]byte(response)))
	doc, _ := goquery.NewDocumentFromReader(bytes.NewReader(response))
	doc.Find("div#user-browser div div div.user-details").Each(func(i int, selection *goquery.Selection) {
		userName := selection.Find("a").Text()
		userLocation := selection.Find("span.user-location").Text()
		if userLocation == "" {
			userLocation = "None"
		}
		var userFlair string
		selection.Find(".-flair > span").Each(func(i int, selection1 *goquery.Selection) {
			temp := selection1.Text()
			userFlair += ", " + temp
		})
		fmt.Println(userName, userLocation, userFlair)
	})
	return nil, nil
}

func (m *Moderators) Single(url string, number int) (interface{}, error) {
	return nil, nil
}
