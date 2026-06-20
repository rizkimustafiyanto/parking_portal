package convert

import (
	"strconv"
	"strings"
)

func ParseUint(rawID string) (uint, error) {
	parsedID, err := strconv.ParseUint(strings.TrimSpace(rawID), 10, 64)
	if err != nil {
		return 0, err
	}

	return uint(parsedID), nil
}

func ParseUintPointer(rawID string) *uint {
	parsedID, err := ParseUint(rawID)
	if err != nil {
		return nil
	}

	return &parsedID
}
