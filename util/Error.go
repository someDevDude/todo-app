package util

//CheckErr checks for error
func CheckErr(err error, handler func(err error)) {
	if err != nil {
		handler(err)
	}
}
