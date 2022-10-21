package main

import (
	"dcadmin/routers"
	"log"
)

func main() {
	log.Printf("系统开始运行！")
	routers.Start()
}
