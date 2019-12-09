package userInterface

import (
	"bytes"
	"github.com/wuxiaoxiaoshen/stackoverflowCli/app/domain/model"
	"github.com/wuxiaoxiaoshen/stackoverflowCli/app/infrastructure"

	"github.com/PuerkitoBio/goquery"
)

type Voters struct {
	model.Votes
}

type VotersResult []Voters

func (v *Voters) FindAll(url string) (interface{}, error) {
	response, _ := infrastructure.ResponseStorager(url)
	//fmt.Println(string([]byte(response)))

	var votersResult VotersResult

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
		if i <= 10 && len(votersResult) <= 10 {
			oneVoters := Voters{
				*model.NewVotes(userName, userFlair, 0),
			}
			votersResult = append(votersResult, oneVoters)
		}

	})
	infrastructure.Marshal(votersResult)
	return nil, nil
}

func (v *Voters) Single(url string, number int) (interface{}, error) {
	return nil, nil
}
