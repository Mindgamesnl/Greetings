package greetings

import (
	"fmt"
	"os"
)

func Write(file string, content string) {
	f, err := os.Create(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = f.WriteString(content)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}
