package uerInterface

import (
	"stackOverFlowClient/app/infrastructure"
	"testing"
)

func TestVoters_FindAll(t *testing.T) {
	tt := []struct {
		url string
	}{
		{
			url: infrastructure.API["voters"],
		},
	}

	var voters Voters
	for _, t := range tt {
		voters.FindAll(t.url)
	}
}
