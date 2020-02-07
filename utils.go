package lifecycle

import "strings"

func escapeID(id string) string {
	return strings.Replace(id, "/", "_", -1)
}
