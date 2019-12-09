package infrastructure

import (
	"encoding/json"
	"fmt"

)

func Marshal(value interface{}) {
	jsonBytes, _ := json.MarshalIndent(value, " ", " ")
	fmt.Println(string(jsonBytes))
}
