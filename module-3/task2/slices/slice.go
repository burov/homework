package slices

import (
	"sort"
)

func SortStringsInPlace(s []string) {
	sort.Strings(s)
}

func SortStringsPure(s []string) []string {
	copy := s
	sort.Strings(copy)
	return copy
}

type User struct {
	FirstName string
	LastName  string
}

type UserList []User

func (u UserList) Len() int {
	return len(u)
}

func (u UserList) Less(i, j int) bool {
	return u[i].FirstName < u[j].FirstName || (u[i].FirstName == u[j].FirstName && u[i].LastName < u[j].LastName)

}

func (u UserList) Swap(i, j int) {
	u[i], u[j] = u[j], u[i]
}

func SortUsersPure(s []User) []User {
	var userList UserList = s
	sort.Sort(userList)
	return userList
}

func SortUsersPureDesc(s []User) []User {
	var userList UserList = s
	sort.Sort(sort.Reverse(userList))
	return userList
}
