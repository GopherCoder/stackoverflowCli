package uerInterface

import (
	"stackOverFlowClient/app/infrastructure"
	"testing"
)

func TestEditors_FindAll(t *testing.T) {
	tt := []struct {
		url string
	}{
		{
			url: infrastructure.API["editors"],
		},
	}
	var editors Editors
	for _, t := range tt {
		editors.FindAll(t.url)
	}
}
