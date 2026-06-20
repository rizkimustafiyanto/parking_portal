package text

import "strings"

func NormalizeStatus(status string) string {
	return strings.ToUpper(strings.TrimSpace(status))
}

func NonEmpty(value string) (string, bool) {
	trimmed := strings.TrimSpace(value)
	if trimmed == "" {
		return "", false
	}

	return trimmed, true
}
