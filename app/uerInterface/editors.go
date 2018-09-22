package uerInterface

import "stackOverFlowClient/app/domain/model"

type Editors struct {
	model.Editor
}

func (e *Editors) FindAll(url string) (interface{}, error) {
	return nil, nil
}

func (e *Editors) Single(url string, number int) (interface{}, error) {
	return nil, nil
}
