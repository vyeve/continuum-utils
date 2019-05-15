package main

import (
	"log"

	"github.com/vitaliyyevenko/continuum-utils/RMM/Kafka-agregate/kafka"
)

func main() {
	end := make(chan bool)
	kafka.Load()
	err := kafka.Load()
	if err != nil {
		log.Fatal(err)
	}
	<-end
}
