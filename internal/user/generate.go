package user

import (
	"fmt"
	"strconv"
	"strings"
)

func GenerateCustomID(id string, prefix string) string {
	number := strings.TrimPrefix(id, prefix)
	lastNumber, _ := strconv.Atoi(number)
	newNumber := lastNumber + 1
	newID := fmt.Sprintf("%s%04d", prefix, newNumber)
	return newID
}
