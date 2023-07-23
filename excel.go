package gocsvexcelize

import (
	"encoding/json"
	"log"
	"strconv"
	"github.com/xuri/excelize/v2"
)

func ConvertExcel(filepath string, hasHeaders bool, sheetName string) string {
	f, err := excelize.OpenFile(filepath)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := f.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	var objects []map[string]string
	var headers []string

	rows, err := f.GetRows(sheetName)

	if hasHeaders {
		headers = rows[0]
	}

	if err != nil {
		log.Fatal(err)
	}
	obj := make(map[string]string)
	for _, row := range rows {
		for i, colCell := range row {
			if hasHeaders {
				headerCell := headers[i]
				obj[headerCell] = colCell
			} else {
				obj[strconv.Itoa(i+1)] = colCell
			}
		}
		objects = append(objects, obj)
	}
	jsonData, err := json.Marshal(objects)
	if err != nil {
		log.Fatal(err)
	}
	return formatBytes(jsonData)
}
