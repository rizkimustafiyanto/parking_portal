package env

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func GetString(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}

	return value
}

func GetInt(key string, defaultValue int) int {
	valueStr := os.Getenv(key)
	if valueStr == "" {
		return defaultValue
	}

	value, err := strconv.Atoi(valueStr)
	if err != nil {
		log.Printf("invalid %s value: %v, defaulting to %d", key, err, defaultValue)
		return defaultValue
	}

	return value
}

func GetBool(key string, defaultValue bool) bool {
	valueStr := os.Getenv(key)
	if valueStr == "" {
		return defaultValue
	}

	value, err := strconv.ParseBool(valueStr)
	if err != nil {
		log.Printf("invalid %s value: %v, defaulting to %t", key, err, defaultValue)
		return defaultValue
	}

	return value
}

func GetStringSlice(
	key string,
	defaultValue []string,
) []string {
	value := os.Getenv(key)

	if value == "" {
		return defaultValue
	}

	parts := strings.Split(value, ",")

	result := make([]string, 0, len(parts))

	for _, part := range parts {
		result = append(
			result,
			strings.TrimSpace(part),
		)
	}

	return result
}