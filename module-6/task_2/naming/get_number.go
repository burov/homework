package naming

import (
	"regexp"
	"strconv"
	"strings"
)

func getNumberSplit(s string) int {
	w := strings.Split(s, `.`)
	if len(w) < 2 {
		return 0
	}
	outString := w[len(w)-1]
	r, err := strconv.ParseInt(outString, 10, 32)
	if err != nil {
		return 0
	}
	return int(r)
}

func getNumberRegexp(s string) int {
	patternLast2Numbers := regexp.MustCompile(`\.([\d]{2})$`)
	if !patternLast2Numbers.MatchString(s) {
		return 0
	}
	d := patternLast2Numbers.FindAllStringSubmatch(s, 1)
	r, err := strconv.ParseInt(d[0][1], 10, 32)
	if err != nil {
		return 0
	}
	return int(r)
}
