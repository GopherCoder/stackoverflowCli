package model

import "reflect"

type Job struct {
	Title   string   `json:"title"`
	Company string   `json:"company"`
	Tags    []string `json:"tags"`
	Salary  string   `json:"salary"`
	Date    string   `json:"date"`
}

type JobResults []Job

func NewJob(title string,
	company string,
	tags []string,
	salary string,
	date string) *Job {
	return &Job{
		Title:   title,
		Company: company,
		Tags:    tags,
		Salary:  salary,
		Date:    date,
	}
}

func (j *Job) JobField(tag string) interface{} {
	value := reflect.ValueOf(&j)
	return value.FieldByName(tag)
}
