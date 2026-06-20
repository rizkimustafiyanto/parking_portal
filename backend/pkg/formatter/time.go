package formatter

import "time"

const DefaultTimeLayout = "2006-01-02 15:04:05"

func FormatTime(value time.Time) string {
	return value.Format(DefaultTimeLayout)
}
