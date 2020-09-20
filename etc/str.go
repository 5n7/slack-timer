package etc

import "regexp"

var spaceRegexp = regexp.MustCompile(`\s+`)

func RemoveDuplicateSpace(s string) string {
	return spaceRegexp.ReplaceAllString(s, " ")
}
