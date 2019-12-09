package userInterface

import (
	"bytes"
	"github.com/wuxiaoxiaoshen/stackoverflowCli/app/domain/model"
	"github.com/wuxiaoxiaoshen/stackoverflowCli/app/infrastructure"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type UserNew struct {
	model.UserNew
}

type UserNewResult []UserNew

func (u *UserNew) FindAll(url string) (interface{}, error) {

	response, _ := infrastructure.ResponseStorager(url)
	//fmt.Println(string([]byte(response)))
	var userNewResult UserNewResult
	doc, _ := goquery.NewDocumentFromReader(bytes.NewReader(response))
	doc.Find("div#user-browser div.grid-layout div div.user-details").Each(func(i int, selection *goquery.Selection) {
		userName := selection.Find("a").Text()
		userLocation := selection.Find("span").Text()
		if userLocation == "" {
			userLocation = "None"
		}

		userFlair := selection.Find("div").Text()
		//fmt.Println(userName, userLocation, strings.TrimSpace(userFlair))
		if i <= 10 && len(userNewResult) <= 10 {
			oneUserNew := UserNew{
				*model.NewUserNew(userName, []string{}, userLocation, strings.TrimSpace(userFlair)),
			}
			userNewResult = append(userNewResult, oneUserNew)
		}

	})
	infrastructure.Marshal(userNewResult)

	return nil, nil
}

func (u *UserNew) Single(url string, number int) (interface{}, error) {
	return nil, nil
}
