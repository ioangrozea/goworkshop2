package main

import (
	"goworkshop2/web"
	"goworkshop2/persistence"
)

func main() {
	if err := persistence.InitDB(); err != nil{
		panic(err)
	}
	web.StartServer()
}

