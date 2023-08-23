package main

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

type Response struct {
	Code int    `json:"code"`
	Data string `json:"data"`
}

func main() {
	nc, err := nats.Connect("nats://nats:4222")

	if err != nil {
		log.Fatalf("Error connecting to NATS: %v", err)
	}
	defer nc.Close()

	_, err = nc.Subscribe("request.echo", func(m *nats.Msg) {
		response := Response{
			Code: 200,
			Data: string(m.Data),
		}
		responseJson, err := jsoniter.Marshal(&response)
		if err != nil {
			nc.Publish("request.echo", []byte(fmt.Sprintf(`{"code":404, "data":"%v"}`, err)))
			time.Sleep(time.Second)
		}
		m.Respond(responseJson)
	})

	if err != nil {
		log.Fatalf("Error subscribing to topic: %v", err)
	}

	for {
		currentTime := time.Now().Format("15:04:05")
		response := Response{
			Code: 200,
			Data: currentTime,
		}
		responseJson, err := jsoniter.Marshal(&response)
		if err != nil {
			nc.Publish("time.tick", []byte(fmt.Sprintf(`{"code":404, "data":"%v"}`, err)))
			time.Sleep(time.Second)
		}
		nc.Publish("time.tick", responseJson)
		fmt.Printf("Published message: %s\n", responseJson)
		time.Sleep(1 * time.Second)
	}
}
