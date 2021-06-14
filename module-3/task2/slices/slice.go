package slices

import (
	"sort"
)

func SortStringsInPlace(s []string) {
	sort.Strings(s)
}

func SortStringsPure(s []string) []string {
	r := make([]string, len(s))
	copy(r, s)

	sort.Strings(r)

	return r
}

type User struct {
	FirstName string
	LastName  string
}

func SortUsersPure(s []User) []User {
	r := make([]User, len(s))

	copy(r, s)

	sort.Slice(r, func(i, j int) bool {
		if r[i].FirstName != r[j].FirstName {
			return r[i].FirstName < r[j].FirstName
		}

		return r[i].LastName < r[j].LastName
	})

	return r
}

func SortUsersPureDesc(s []User) []User {
	r := make([]User, len(s))

	copy(r, s)

	sort.Slice(r, func(i, j int) bool {
		if r[i].FirstName != r[j].FirstName {
			return r[i].FirstName > r[j].FirstName
		}

		return r[i].LastName > r[j].LastName
	})

	return r
}
