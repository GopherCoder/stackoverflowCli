package uerInterface

import (
	"github.com/wuxiaoxiaoshen/stackoverflowCli/app/infrastructure"
	"testing"
)

func TestJob_FindAll(t *testing.T) {
	tt := []struct {
		url string
	}{
		{
			url: infrastructure.API["jobs"],
		},
	}

	var job Job
	for _, t := range tt {
		job.FindAll(t.url)
	}
}
