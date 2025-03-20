package nats

import (
	"log"

	"github.com/dvl-numeez/nats-demo/config"
	"github.com/nats-io/nats.go"
)

func JetStreamInit() (nats.JetStreamContext, error) {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		return nil, err
	}
	js, err := nc.JetStream(nats.PublishAsyncMaxPending(256))
	if err != nil {
		return nil, err
	}
	err = CreateStream(js)
	if err != nil {
		return nil, err
	}

	return js, nil
}

func CreateStream(jetStream nats.JetStreamContext) error {
	stream, _ := jetStream.StreamInfo(config.StreamName)
	if stream == nil {
		log.Printf("Creating stream: %s\n", config.StreamName)

		_, err := jetStream.AddStream(&nats.StreamConfig{
			Name:     config.StreamName,
			Subjects: []string{config.StreamSubjects},
		})
		if err != nil {
			return err
		}
	}
	return nil
}
