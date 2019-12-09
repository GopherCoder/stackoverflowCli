package uerInterface

import (
	"bytes"
	"github.com/wuxiaoxiaoshen/stackoverflowCli/app/domain/model"
	"github.com/wuxiaoxiaoshen/stackoverflowCli/app/infrastructure"

	"github.com/PuerkitoBio/goquery"
)

type Editors struct {
	model.Editor
}

type EditorResult []Editors

func (e *Editors) FindAll(url string) (interface{}, error) {
	response, _ := infrastructure.ResponseStorager(url)
	//fmt.Println(string([]byte(response)))

	var editorResult EditorResult
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
		OneEditor := Editors{
			*model.NewEditor(
				userName, []string{}, userLocation, userFlair, 0),
		}
		if i < 10 {
			editorResult = append(editorResult, OneEditor)
		}

	})
	infrastructure.Marshal(editorResult)
	return nil, nil
}

func (e *Editors) Single(url string, number int) (interface{}, error) {
	return nil, nil
}
