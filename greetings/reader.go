package greetings

import (
	"fmt"
	"io/ioutil"
)

func Read(file string) string {
	b, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Print(err)
	}
	return string(b)
}
