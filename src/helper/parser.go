package helper

import (
	"fmt"
	"strconv"
)

func ConvertString(str string) int {
	id, error := strconv.Atoi(str)
	if error != nil {
		fmt.Println("Error converting string")
		return 0
	}
	return id
}
