package uerInterface

import (
	"bytes"
	"stackOverFlowClient/app/domain/model"
	"stackOverFlowClient/app/infrastructure"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Job struct {
	model.Job
}

type JobResult []Job

func (j *Job) FindAll(url string) (interface{}, error) {
	response, _ := infrastructure.ResponseStorager(url)

	//fmt.Println(string([]byte(response)))

	var jobResult JobResult
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
		//fmt.Println("title", jobTitle, "date", jobDate, "description", jobDescription, "salary", jobSalary, "tags", jobTags)
		if i < 10 || len(jobResult) >= 10 {
			oneJob := Job{
				*model.NewJob(jobTitle,
					jobDate,
					jobTags,
					jobSalary,
					jobDescription),
			}
			jobResult = append(jobResult, oneJob)

		}

	})
	infrastructure.Marshal(jobResult)
	return nil, nil
}

func (j *Job) Single(url string, number int) (interface{}, error) {
	return nil, nil
}
