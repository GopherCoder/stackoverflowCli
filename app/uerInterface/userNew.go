package uerInterface

import "stackOverFlowClient/app/domain/model"

type UserNew struct {
	model.UserNew
}

func (u *UserNew) FindAll(url string) (interface{}, error) {
	return nil, nil
}

func (u *UserNew) Single(url string, number int) (interface{}, error) {
	return nil, nil
}
