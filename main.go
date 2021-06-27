package main

import (
	"flag"
	"fmt"

	"github.com/coala/corobo-ng/config"
	"github.com/coala/corobo-ng/db"
	"github.com/coala/corobo-ng/server"
)

func main() {
	environment := flag.String("e", "development", "Mode in which server should run")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
	}
	flag.Parse()
	config.Init(*environment)
	db.Init()
	server.Init()
}
