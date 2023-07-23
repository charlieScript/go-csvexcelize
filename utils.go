package gocsvexcelize

import (
	"bytes"
	"encoding/json"
	"log"
)

func formatBytes(data []byte) string {
	var formattedJSON bytes.Buffer
	err := json.Indent(&formattedJSON, []byte(data), "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return formattedJSON.String()
}