package naming

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const maxNumber = 99

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

func getNumberFromName(s string) int {
	return getNumberSplit(s)
}

func nameWithNumber(name string, n int) string {
	if n == 0 {
		return name
	}
	return fmt.Sprintf("%s.%02d", name, n)
}

func NameWithNumberIncrement(baseName string, names []string) (string, error) {
	nMax := -1
	for _, nm := range names {
		if n := getNumberFromName(nm); n > nMax {
			nMax = n
		}
	}
	nMax++
	if nMax > maxNumber {
		return "", ErrMaxNamesCountHit
	}
	return nameWithNumber(baseName, nMax), nil
}
