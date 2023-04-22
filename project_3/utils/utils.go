package utils

import "fmt"

func CkeckError(err error) {
	if err != nil {
		fmt.Println(err.Error())

	}
}
