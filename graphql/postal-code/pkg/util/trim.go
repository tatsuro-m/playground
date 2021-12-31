package util

import "strings"

func TrimFullWidthSpace(s string) string {
	return strings.Replace(s, "　", "", -1)
}
