package helpers

import (
	"fmt"
)

func PanicIfError(err error, message string) {
	if err != nil {
		panic(message)
	}
}

func LogError(err error) {
	if err != nil {
		// do things
		fmt.Println(err.Error())
	}
}
