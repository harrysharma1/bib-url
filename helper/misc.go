package helper

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

func FormatMiscBibtex(
	miscCiteKey string,
	miscFields []string,
	braces bool) string {
	var sb strings.Builder
	sb.WriteString("@misc{")
	if miscCiteKey != "" {
		sb.WriteString(miscCiteKey)
	} else {
		sb.WriteString(uuid.NewString())
	}
	sb.WriteString(",\n")
	customFieldMap := parseFields(miscFields)
	fields := []string{}
	wrap := func(field string, s string) string {
		if braces {
			return "{" + s + "}"
		}

		if slices.Contains(months, s) {
			return s
		}
		if field == "year" {
			if _, err := strconv.Atoi(s); err == nil {
				return s
			}
		}

		return `"` + s + `"`

	}

	// Largest key
	keys := []string{}
	for key := range customFieldMap {
		keys = append(keys, key)
	}
	maxKeyLength := largestFieldInt(keys)

	for key, val := range customFieldMap {
		var tmpValue string
		switch key {
		case "author":
			tmpValue = wrap(key, strings.Join(val, " and "))
		case "year":
			// Decided to pick first only
			tmpValue = wrap(key, val[0])
		default:
			tmpValue = wrap(key, strings.Join(val, " "))
		}
		fields = append(fields,
			fmt.Sprintf("\t%-*s = %s", maxKeyLength, key, tmpValue),
		)
	}

	sb.WriteString(strings.Join(fields, ",\n"))
	sb.WriteString("\n}")
	return sb.String()
}
