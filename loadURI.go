package main

import "github.com/mubaris/graviton"

func main() {
	github := graviton.NewApp("Github", 1200, 800)
	driver := graviton.NewDriver(*github)
	driver.Initialize()
	driver.AttachURI("https://github.com")
	driver.Start()
}
