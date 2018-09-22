package uerInterface

import "stackOverFlowClient/app/domain/model"

type Moderators struct {
	model.Moderators
}

func (m *Moderators) FindAll(url string) (interface{}, error) {
	return nil, nil
}

func (m *Moderators) Single(url string, number int) (interface{}, error) {
	return nil, nil
}
