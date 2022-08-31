package utils

import (
	"encoding/json"
	"fmt"
	"log"
)

func PrettyPrint(data interface{}) {
	val, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		log.Fatalf("error printing data %s", err)
	}

	fmt.Println(string(val))
}
