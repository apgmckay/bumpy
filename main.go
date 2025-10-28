package main

import (
	"bumpy/cmd"
	"fmt"
	"os"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		errorExit(1, err.Error())
	}
}

func errorExit(xCode int, xMsg string) {
	fmt.Println(xMsg)
	os.Exit(xCode)
}
