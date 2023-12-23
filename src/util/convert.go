package util

import (
	"encoding/json"
	"fmt"
)

func Convert(result any) (string, error) {

	jsonData, err := json.Marshal(result)
	if err != nil {
		fmt.Println(err)
	}

	// Convertendo os dados JSON para string
	jsonString := string(jsonData)

	return jsonString, err
}
