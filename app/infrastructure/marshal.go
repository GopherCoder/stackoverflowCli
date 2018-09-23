package infrastructure

import (
	"fmt"

	"github.com/gin-gonic/gin/json"
)

func Marshal(value interface{}) {
	jsonBytes, _ := json.MarshalIndent(value, " ", " ")
	fmt.Println(string(jsonBytes))
}
