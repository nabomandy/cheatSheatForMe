package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// replace this by fetching actual response body
	responseBody := `{"column_name": ["geom","sublayer_id","install_date","address","direction","distance","emerg_bell","emerg_bell_board"]}`
	var data map[string]interface{}
	err := json.Unmarshal([]byte(responseBody), &data)
	if err != nil {
		panic(err)
	}
	fmt.Println(data["column_name"])
}
