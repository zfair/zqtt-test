package test

import (
	Mqtt "github.com/eclipse/paho.mqtt.golang"
	"testing"
)

const endpoint = "tcp://127.0.0.1:1883"

func TestConnect(t *testing.T) {
	opts := Mqtt.NewClientOptions()
	opts.AddBroker(endpoint)
	opts.SetClientID("0")
	opts.SetUsername("test")
	opts.SetPassword("test")

	client := Mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		t.Fatal(token.Error())
	}
}
