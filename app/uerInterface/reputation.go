package uerInterface

import "stackOverFlowClient/app/domain/model"

type Reputation struct {
	model.Reputation
}

type Reputations []Reputation

func (r *Reputation) FindAll(url string) (interface{}, error) {
	return Reputations{}, nil
}

func (r *Reputation) Single(url string, number int) (interface{}, error) {
	return Reputation{}, nil
}
