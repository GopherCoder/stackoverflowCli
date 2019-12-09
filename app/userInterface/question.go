package userInterface

import (
	"bytes"
	"fmt"
	"github.com/wuxiaoxiaoshen/stackoverflowCli/app/domain/model"
	"github.com/wuxiaoxiaoshen/stackoverflowCli/app/infrastructure"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Question struct {
	model.Question
}

type QuestionResult []Question

func (q *Question) FindAll(url string) (interface{}, error) {
	response, _ := infrastructure.ResponseStorager(url)
	//fmt.Println(string([]byte(response)))

	var questionResult QuestionResult
	doc, _ := goquery.NewDocumentFromReader(bytes.NewReader(response))
	doc.Find("div.flush-left.js-search-results div.question-summary.search-result").Each(func(i int, selection *goquery.Selection) {
		votes := strings.TrimSpace(selection.Find(".statscontainer .stats").Find("div.vote").Text())
		newVotes := strings.Replace(strings.Replace(votes, "\n", "", -1), " ", "", -1)
		summary := selection.Find(".summary")
		title, _ := summary.Find(".result-link h3 a").Attr("title")
		url, _ := summary.Find(".result-link h3 a").Attr("href")
		fullUrl := "https://stackoverflow.com" + url
		description := strings.Replace(strings.TrimSpace(summary.Find(".excerpt").Text()), "\n", "", -1)
		information := strings.TrimSpace(summary.Find(".started.fr").Text())
		var questionType string
		if strings.Contains(information, "answered") {
			questionType = "answer"
		} else {
			questionType = "ask"
		}
		fmt.Println(newVotes, title, description, information, questionType, fullUrl)
		newVotesInt, _ := strconv.Atoi(newVotes)
		if i <= 10 && len(questionResult) <= 10 {
			oneQuestion := Question{
				*model.NewQuestion(newVotesInt, questionType, title, fullUrl, description, information),
			}
			questionResult = append(questionResult, oneQuestion)
		}

	})
	infrastructure.Marshal(questionResult)
	return nil, nil
}

func (q *Question) Single(url string, number int) (interface{}, error) {
	return nil, nil
}
