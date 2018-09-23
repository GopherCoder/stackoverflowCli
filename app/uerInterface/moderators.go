package uerInterface

import (
	"bytes"
	"stackOverFlowClient/app/domain/model"
	"stackOverFlowClient/app/infrastructure"

	"github.com/PuerkitoBio/goquery"
)

type Moderators struct {
	model.Moderators
}

type ModeratorsResult []Moderators

func (m *Moderators) FindAll(url string) (interface{}, error) {
	response, _ := infrastructure.ResponseStorager(url)
	//fmt.Println(string([]byte(response)))

	var moderatorsResult ModeratorsResult

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
		//fmt.Println(userName, userLocation, userFlair)
		if i < 10 && len(moderatorsResult) <= 10 {
			oneModerators := Moderators{
				*model.NewModerators(userName, []string{}, userLocation, userFlair, ""),
			}
			moderatorsResult = append(moderatorsResult, oneModerators)
		}

	})
	infrastructure.Marshal(moderatorsResult)
	return nil, nil
}

func (m *Moderators) Single(url string, number int) (interface{}, error) {
	return nil, nil
}
