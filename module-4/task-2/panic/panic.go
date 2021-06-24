package panic

func iWillPanic() {
	panic("something")
}

func catchPanic() {
	iWillPanic()
}
