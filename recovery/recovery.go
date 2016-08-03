package recovery

func Recover(callback func(error)) {
	if r := recover(); r != nil {
		switch err := r.(type) {
		case error:
			callback(err)
		}
	}
}
