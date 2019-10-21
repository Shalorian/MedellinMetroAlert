package main

import (
	"log"

	"github.com/MedellinMetroAlert/supervisor"
	"github.com/MedellinMetroAlert/utils"
)

func main() {
	log.Println("Initializing configuration ...")
	utils.InitializeConfigUtils()

	log.Println("Starting supervisor ...")
	supervisor.Start()
}
