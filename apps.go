package main

import (
	"os"

	"sflowbeat/beater"

	"github.com/elastic/beats/libbeat/beat"
)

var Name = "flowbeat"

func main() {
	if err := beat.Run(Name, "", beater.New()); err != nil {
		os.Exit(1)
	}
}
