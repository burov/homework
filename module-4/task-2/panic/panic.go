package panic

func iWillPanic() {
	panic("something")
}

func catchPanic() {
	defer func() {
		if err := recover(); err != nil {
			// No panic, everything is under control.
		}
	}()
	iWillPanic()
}
