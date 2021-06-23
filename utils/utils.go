package utils

import "log"

// HandleError function for handle error
func HandleError(err error) {
	if err != nil {
		log.Panic(err)
	}
}
