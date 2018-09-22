package model

import "reflect"

type Tag struct {
	Name        string `json:"name"`
	Number      int    `json:"number"`
	Description string `json:"description"`
	Info        string `json:"info"`
}

type TagResults []Tag

func NewTag(
	name string,
	number int,
	description string,
	info string) *Tag {
	return &Tag{
		Name:        name,
		Number:      number,
		Description: description,
		Info:        info,
	}
}

func (t *Tag) TagField(tag string) interface{} {
	value := reflect.ValueOf(&t)
	return value.FieldByName(tag)
}
