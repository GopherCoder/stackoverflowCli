package uerInterface

import "stackOverFlowClient/app/domain/model"

type Voters struct {
	model.Votes
}

func (v *Voters) FindAll(url string) (interface{}, error) {
	return nil, nil
}

func (v *Voters) Single(url string, number int) (interface{}, error) {
	return nil, nil
}
