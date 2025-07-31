package main

import (
	"bumpy/cmd"
	"log"
	"os"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}
}
