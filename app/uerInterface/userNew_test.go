package uerInterface

import (
	"github.com/wuxiaoxiaoshen/stackoverflowCli/app/infrastructure"
	"testing"
)

func TestUserNew_FindAll(t *testing.T) {
	tt := []struct {
		url string
	}{
		{
			url: infrastructure.API["newUser"],
		},
	}

	var userNew UserNew
	for _, t := range tt {
		userNew.FindAll(t.url)
	}
}
