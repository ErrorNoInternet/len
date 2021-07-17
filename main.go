package main

import (
	"fmt"
	"io"
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
			file, errorObject := os.Open(argument)
			if errorObject != nil {
				fmt.Println(errorObject)
				os.Exit(1)
			}
			buffer := make([]byte, 8000000)
			for {
				size, errorObject := file.Read(buffer)
				length += size
				if errorObject == io.EOF {
					break
				} else if errorObject != nil {
					fmt.Println(errorObject)
					os.Exit(1)
				}
			}
		}
	}
	fmt.Println(length)
}
