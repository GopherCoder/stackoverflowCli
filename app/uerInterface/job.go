package uerInterface

import (
	"bytes"
	"fmt"
	"stackOverFlowClient/app/domain/model"
	"stackOverFlowClient/app/infrastructure"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Job struct {
	model.Job
}

func (j *Job) FindAll(url string) (interface{}, error) {
	response, _ := infrastructure.ResponseStorager(url)

	//fmt.Println(string([]byte(response)))

	doc, _ := goquery.NewDocumentFromReader(bytes.NewReader(response))

	doc.Find("div.listResults div div.-job-summary").Each(func(i int, selection *goquery.Selection) {
		jobTitle := selection.Find("div h2  a").Text()
		jobDate := selection.Find("div.-title span").Last().Text()
		var jobDescription string
		selection.Find("div").Eq(1).Find("span").Each(func(i int, selection1 *goquery.Selection) {
			temp := strings.TrimSpace(selection1.Text())
			newTemp := strings.Replace(temp, "\n", "", -1)
			jobDescription += " " + newTemp

		})

		var jobSalary string
		selection.Find("div").Eq(2).Find("span").Each(func(i int, selection2 *goquery.Selection) {

			temp := strings.TrimSpace(selection2.Text())
			newTemp := strings.Replace(temp, "\n", "", -1)
			jobSalary += " " + newTemp

		})

		var jobTags string
		selection.Find("div").Last().Find("a").Each(func(i int, selection3 *goquery.Selection) {
			jobTags += " " + strings.TrimSpace(selection3.Text())

		})
		fmt.Println("title", jobTitle, "date", jobDate, "description", jobDescription, "salary", jobSalary, "tags", jobTags)
	})

	return nil, nil
}

func (j *Job) Single(url string, number int) (interface{}, error) {
	return nil, nil
}
