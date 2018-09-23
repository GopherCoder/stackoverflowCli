package uerInterface

import (
	"bytes"
	"fmt"
	"stackOverFlowClient/app/domain/model"
	"stackOverFlowClient/app/infrastructure"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type UserNew struct {
	model.UserNew
}

func (u *UserNew) FindAll(url string) (interface{}, error) {

	response, _ := infrastructure.ResponseStorager(url)
	fmt.Println(string([]byte(response)))
	doc, _ := goquery.NewDocumentFromReader(bytes.NewReader(response))
	doc.Find("div#user-browser div.grid-layout div div.user-details").Each(func(i int, selection *goquery.Selection) {
		fmt.Println(123, i)
		userName := selection.Find("a").Text()
		userLocation := selection.Find("span").Text()
		if userLocation == "" {
			userLocation = "None"
		}

		userFlair := selection.Find("div").Text()
		fmt.Println(userName, userLocation, strings.TrimSpace(userFlair))

	})

	return nil, nil
}

func (u *UserNew) Single(url string, number int) (interface{}, error) {
	return nil, nil
}
