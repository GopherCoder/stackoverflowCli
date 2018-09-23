package uerInterface

import (
	"fmt"
	"stackOverFlowClient/app/infrastructure"
	"testing"
)

func TestQuestion_FindAll(t *testing.T) {
	tt := []struct {
		url string
	}{
		{
			url: infrastructure.API["search"],
		},
	}
	var question Question
	for _, t := range tt {
		question.FindAll(fmt.Sprintf(t.url, "go"))
	}
}
