package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	var length int
	buffer := make([]byte, 16000000)

	if len(os.Args) == 1 {
		for {
			size, errorObject := os.Stdin.Read(buffer)
			length += size
			if errorObject == io.EOF {
				break
			} else if errorObject != nil {
				fmt.Println(errorObject)
				os.Exit(1)
			}
		}
	} else {
		os.Args = os.Args[1:]
		for _, argument := range os.Args {
			file, errorObject := os.Open(argument)
			if errorObject != nil {
				fmt.Println(errorObject)
				os.Exit(1)
			}
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
