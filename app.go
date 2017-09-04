package main

import "github.com/astaxie/beego"

// Controller ...
type Controller struct {
	beego.Controller
}

// Get controller ...
func (c *Controller) Get() {
	c.Ctx.WriteString("Wassup bruh")
}

func main() {
	// define the router, and the controller ...
	beego.Router("/", &Controller{})

	// running in port 8080
	beego.Run()
}
