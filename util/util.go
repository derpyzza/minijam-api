package util

import "log"

func CheckErr(err error, errorMsg ...string) {

	// FUIIYOOOO
	var msg string

	if len(errorMsg) > 0 {
		msg = errorMsg[0]
	} else {
		msg = "ERROR: "
	}
	if err != nil {
		log.Fatalln(msg, err)
	}
}
