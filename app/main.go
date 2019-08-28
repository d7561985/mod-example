package main

import (
	"flag"
	"something/app/pkg/filter"
	"something/app/pkg/logger"
)

func main() {
	str := ""
	flag.StringVar(&str, "str", "empty", "upper case")
	flag.Parse()

	logger.Fatal(filter.UP(str))
}