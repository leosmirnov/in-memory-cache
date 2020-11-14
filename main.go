package main

import (
	"flag"
)

func main() {

	cfgPath := flag.String("config", "./config.yml", "full file path to config")
	flag.Parse()

}
