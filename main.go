package main

import (
	"bumpy/cmd"
	"fmt"
	"os"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
