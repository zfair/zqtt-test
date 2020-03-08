package test

import (
	"testing"
	MQTT "github.com/eclipse/paho.mqtt.golang"
)

const endpoint = "tcp://127.0.0.1:12345"

func TestConnect(t *testing.T) {
	opts := MQTT.NewClientOptions()
	opts.AddBroker(endpoint)
	opts.SetClientID("0")
	opts.SetUsername("test")
	opts.SetPassword("test")

	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		t.Fatal(token.Error())
	}
}