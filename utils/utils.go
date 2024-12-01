package utils

import (
	"regexp"
	"strconv"
)

func ParseInt(input string) int {
	re := regexp.MustCompile("[0-9]+")
	idString := re.FindString(input)
	id, err := strconv.Atoi(idString)

	if err != nil {
		panic(err)
	}

	return id
}
