package infrastructure

import (
	"fmt"
	"testing"
)

func TestResponseStorager(t *testing.T) {
	tt := []struct {
		url string
	}{
		{
			url: "http://www.baidu.com",
		}, {
			url: "https://stackoverflow.com/users?tab=moderators",
		},
	}
	for _, t := range tt {
		byteResponse, _ := ResponseStorager(t.url)
		fmt.Println(string(byteResponse))
	}
}
