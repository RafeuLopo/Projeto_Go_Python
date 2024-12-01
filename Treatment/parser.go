package treatment

import (
	"strconv"
	"strings"
	"time"
)

func parseDate(dateStr string) *time.Time {
	dateStr = strings.TrimSpace(dateStr)
	if dateStr == "" {
		return nil
	}
	t, err := time.Parse("2006/01/02 15:04:05", dateStr)
	if err != nil {
		return nil
	}
	return &t
}

func parseFloat64(val string) *float64 {
	val = strings.TrimSpace(val)
	if val == "" {
		return nil
	}
	val = strings.ReplaceAll(val, ",", ".") // Normaliza números com vírgulas
	parsed, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return nil
	}
	return &parsed
}

func parseInt(val string) int {
	val = strings.TrimSpace(val)
	parsed, err := strconv.Atoi(val)
	if err != nil {
		return 0
	}
	return parsed
}
