package uerInterface

import (
	"bytes"
	"fmt"
	"stackOverFlowClient/app/domain/model"
	"stackOverFlowClient/app/infrastructure"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Tags struct {
	model.Tag
}

func (t *Tags) FindAll(url string) (interface{}, error) {
	response, _ := infrastructure.ResponseStorager(url)
	fmt.Println(string([]byte(response)))
	doc, _ := goquery.NewDocumentFromReader(bytes.NewReader(response))
	doc.Find("div#tags-browser div.grid-layout--cell.tag-cell").Each(func(i int, selection *goquery.Selection) {
		tagName := selection.Find("a.post-tag").Text()
		tagNumber := selection.Find("span.item-multiplier-count").Text()
		tagDescription := selection.Find("div.excerpt").Text()

		var tagInfo string
		selection.Find("div.grid.jc-space-between div.grid--cell.stats-row a").Each(func(i int, selection1 *goquery.Selection) {
			tagInfo += " " + strings.TrimSpace(selection1.Text())
		})

		fmt.Println(tagName, tagNumber, tagDescription, tagInfo)

	})
	return nil, nil
}
func (t *Tags) Single(url string, number int) (interface{}, error) {
	return nil, nil
}
