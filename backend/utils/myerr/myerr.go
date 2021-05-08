package myerr

import "log"

// CheckErr checking err
func CheckErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
