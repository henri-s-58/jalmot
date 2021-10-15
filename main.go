package main

import (
	"log"
	"os"

	"github.com/jalmot/v1"
)

func main() {
	var err error = jalmot.New("jalmot is super error")
	log.Println("err:", err)
	log.Println("err location:", jalmot.TakeLocation(err))
	os.Exit(0)
}
