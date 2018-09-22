package uerInterface

import (
	"stackOverFlowClient/app/infrastructure"
	"testing"
)

func TestReputation_FindAll(t *testing.T) {
	tt := []struct {
		url string
	}{
		{
			url: infrastructure.API["reputation"],
		},
	}

	var reputation Reputation
	for _, t := range tt {
		reputation.FindAll(t.url)

	}
}
