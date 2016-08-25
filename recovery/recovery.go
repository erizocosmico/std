package recovery

// Recover recovers and if it's an error passes it to the given callback. The
// call to this function needs to be deferred.
//   defer Recover(func(err error) {
//       // stuff
//   })
func Recover(callback func(error)) {
	if r := recover(); r != nil {
		switch err := r.(type) {
		case error:
			callback(err)
		}
	}
}
