package maps

import (
	"fmt"
	"sort"
	"strings"
)

func SortedValues(m map[int]string) string {
	var keys []int
	for key := range m {
		keys = append(keys, key)
	}
	sort.Ints(keys)
	fmt.Printf("Keys: %v\n", keys)

	var sb strings.Builder
	for _, key := range keys {
		val := m[key]
		sb.WriteString(val)
	}
	return sb.String()
}
