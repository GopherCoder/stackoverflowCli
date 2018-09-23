package uerInterface

import (
	"stackOverFlowClient/app/infrastructure"
	"testing"
)

func TestModerators_FindAll(t *testing.T) {
	tt := []struct {
		url string
	}{
		{
			url: infrastructure.API["moderators"],
		},
	}
	var moderators Moderators
	for _, t := range tt {
		moderators.FindAll(t.url)

	}
}
