package main

import (
	"kalendar/lib"
	"kalendar/server"
)

func main() {
	server.StartServer(lib.GetEnv("PORT", ":9999"))
}
