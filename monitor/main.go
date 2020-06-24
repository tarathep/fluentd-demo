package main

import (
	"os"

	"github.com/tarathep/kafka"
)

func main() {
	//dev.tarathep.com:9092
	//log-messages
	kafka.Subscribe(string(os.Getenv("BROKER")), os.Getenv("TOPIC"))
}
