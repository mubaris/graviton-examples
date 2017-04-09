package main

import "github.com/mubaris/graviton"

func main() {
	test := graviton.NewApp("Test", 600, 400)
	driver := graviton.NewDriver(*test)
	driver.Initialize()
	driver.AttachHTML("<p>It Works</p>", "")
	driver.Start()
}
