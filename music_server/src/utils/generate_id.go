package utils

import (
	"github.com/google/uuid"
	"strings"
)

func GenerateID() string {
	strId := uuid.New().String()
	newStrId := strings.ReplaceAll(strId, "-", "")
	return newStrId
}
