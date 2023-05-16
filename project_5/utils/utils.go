package utils

import (
	"fmt"
	"io/ioutil"
)

func CkeckError(err error) {
	if err != nil {
		fmt.Println("Error: ", err.Error())
	}
}

func LoadFile(fileName string) (string, error) {
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		CkeckError(err)
		return "", err
	}
	return string(bytes), nil
}
