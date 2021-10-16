package main

import (
	"log"
	"os"

	"github.com/henri-s-58/jalmot/v1"
)

func main() {
	var err error = jalmot.Newf("jalmot is super %s. made by: %s", "error", "henri-s-58")
	log.Println("err:", err)
	log.Println("err location:", jalmot.TakeLocation(err))
	os.Exit(0)
}
