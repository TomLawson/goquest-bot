package main

// CheckErr will panic if an error is not nil. For the time being I'll use it as a quick and dirty way to stay DRY...
func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
