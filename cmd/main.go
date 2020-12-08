package main

import (
	"github.com/fanama/goSqLite/kit/console"
	"github.com/fanama/goSqLite/kit/tools"
	"github.com/fanama/goSqLite/platform/greetings"
)

func main() {
	tools.InitValues()
	greetings.Helloworld()
	console.Log(tools.Hello + " " + tools.Name)
	console.Timer(tools.Time)
}
