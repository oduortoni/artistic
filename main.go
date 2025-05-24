package main

import (
	_ "artistic/routers"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
    beego.SetStaticPath("/", "frontend/dist")
	beego.Run()
}
