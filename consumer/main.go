package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	n "github.com/dvl-numeez/nats-demo"
	"github.com/dvl-numeez/nats-demo/config"
	"github.com/nats-io/nats.go"
)

func main() {

	js, err := n.JetStreamInit()
	if err != nil {
		log.Fatal(err)
	}

	_, err = js.Subscribe(config.SubjectName, func(m *nats.Msg) {

		err = m.Ack()
		if err != nil {
			log.Println("Unable to Ack", err)
			return
		}
		fmt.Printf("Message from publisher: %s\n", string(m.Data))
	},
		nats.ManualAck(),
		nats.DeliverNew(),
		nats.Durable(config.DurableName))

	if err != nil {
		log.Println("Subscribe failed")
		return
	}

	quitChannel := make(chan os.Signal, 1)
	signal.Notify(quitChannel, syscall.SIGINT, syscall.SIGTERM)
	<-quitChannel
	log.Println("\nShutting consumer")

}
