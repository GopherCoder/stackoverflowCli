package uerInterface

import (
	"stackOverFlowClient/app/infrastructure"
	"testing"
)

func TestTags_FindAll(t *testing.T) {
	tt := []struct {
		url string
	}{
		{
			url: infrastructure.API["tags"],
		},
	}

	var tags Tags
	for _, t := range tt {
		tags.FindAll(t.url)
	}
}
