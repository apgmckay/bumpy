package main

import "bumpy/package/server"

func main() {
	bumpy := server.New()
	bumpy.Run()
}
