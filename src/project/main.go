package main

import (
	_ "project/routers"
	_ "project/sysinit"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}

