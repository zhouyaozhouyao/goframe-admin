package main

import (
	_ "gadmin/boot"
	_ "gadmin/router"

	"github.com/gogf/gf/frame/g"
)

func main() {
	g.Server().Run()
}
