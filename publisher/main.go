package main

import (
	"fmt"
	"log"
	"time"

	nats "github.com/dvl-numeez/nats-demo"
	"github.com/dvl-numeez/nats-demo/config"
)

func main() {
	js, err := nats.JetStreamInit()
	if err != nil {
		log.Fatal("JS Stream Init error : ", err)
	}

	for i := 0; i <= 100; i++ {
		fmt.Printf("Publishing message %d\n", i)
		message := fmt.Sprintf("message %d", i)
		_, err = js.Publish(config.SubjectName, []byte(message))
		time.Sleep(2 * time.Second)
		if err != nil {
			log.Println(err)
		}
	}

}
