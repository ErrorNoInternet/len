package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	var length int

	if len(os.Args) == 1 {
		data, errorObject := ioutil.ReadAll(os.Stdin)
		if errorObject != nil {
			fmt.Println(errorObject)
			os.Exit(1)
		}
		length += len(data)
	} else {
		os.Args = os.Args[1:]
		for _, argument := range os.Args {
			data, errorObject := ioutil.ReadFile(argument)
			if errorObject != nil {
				fmt.Println(errorObject)
				os.Exit(1)
			}
			length += len(data)
		}
	}
	fmt.Println(length)
}
