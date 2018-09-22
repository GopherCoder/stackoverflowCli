package model

import "reflect"

type Question struct {
	Votes       int    `json:"votes"`
	Type        string `json:"type"`
	Title       string `json:"title"`
	URL         string `json:"url"`
	Description string `json:"description"`
	Information string `json:"information"`
}

type QuestionResults []Question

func NewQuestion(
	votes int,
	typeInfo string,
	title string,
	url string,
	description string,
	information string) *Question {
	return &Question{
		Votes:       votes,
		Type:        typeInfo,
		Title:       title,
		URL:         url,
		Description: description,
		Information: information,
	}
}

func (q *Question) QuestionField(tag string) interface{} {
	value := reflect.ValueOf(&q)
	return value.FieldByName(tag)
}
