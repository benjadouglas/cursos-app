package utils

import "log"

func FailOnErr(err error, message string) {
	if err != nil {
		log.Panicf("%v %v", message, err)
	}
}
