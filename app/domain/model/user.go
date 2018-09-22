package model

import "reflect"

type commonTagField struct {
	Name string   `json:"name"`
	Tags []string `json:"tags"`
}

// Reputation
type Reputation struct {
	commonTagField
	Address string `json:"address"`
	Flair   string `json:"flair"`
}

func NewReputation(
	name string,
	tags []string,
	address string,
	Flair string) *Reputation {
	return &Reputation{
		commonTagField{Name: name, Tags: tags},
		address,
		Flair,
	}
}

func (r *Reputation) ReputationField(tag string) interface{} {
	value := reflect.ValueOf(&r)
	return value.FieldByName(tag)
}

type UserNew struct {
	commonTagField
	Address string `json:"address"`
	Flair   string `json:"flair"`
}

func NewUserNew(
	name string,
	tags []string,
	address string,
	flair string) *UserNew {
	return &UserNew{
		commonTagField{Name: name, Tags: tags},
		address,
		flair,
	}
}

func (u *UserNew) UserNewField(tag string) interface{} {
	value := reflect.ValueOf(&u)
	return value.FieldByName(tag)
}

type Votes struct {
	Name   string `json:"name"`
	Flair  string `json:"flair"`
	Number int    `json:"number"`
}

func NewVotes(
	name string,
	flair string,
	number int) *Votes {
	return &Votes{
		Name:   name,
		Flair:  flair,
		Number: number,
	}
}

func (v *Votes) VotesField(tag string) interface{} {
	value := reflect.ValueOf(&v)
	return value.FieldByName(tag)
}

type Editor struct {
	Reputation
	Number int `json:"number"`
}

func NewEditor(
	name string,
	tags []string,
	address string,
	flair string,
	number int) *Editor {
	return &Editor{
		Reputation{
			commonTagField{Name: name, Tags: tags},
			address,
			flair,
		},
		number,
	}
}

func (e *Editor) EditorField(tag string) interface{} {
	value := reflect.ValueOf(&e)
	return value.FieldByName(tag)
}

type Moderators struct {
	Reputation
	ElectedTime string `json:"elected_time"`
}

func NewModerators(
	name string,
	tags []string,
	address string,
	flair string,
	electedTime string) *Moderators {
	return &Moderators{Reputation{
		commonTagField{Name: name, Tags: tags},
		address,
		flair,
	},
		electedTime}
}

func (m *Moderators) ModeratorsField(tag string) interface{} {
	value := reflect.ValueOf(&m)
	return value.FieldByName(tag)
}
