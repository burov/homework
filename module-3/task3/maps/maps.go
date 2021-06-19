package maps

import (
	"sort"
	"strings"
)

func SortedValues(m map[int]string) string {
	keys := make([]int, 0, len(m))
	result := make([]string, 0, len(m))

	for k := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for _, v := range keys {
		result = append(result, m[v])
	}

	return strings.Join(result, "")
}
