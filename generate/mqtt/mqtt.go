package mqtt

import (
	"fmt"
	"log"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

var BrokerHost string
var Port = 1883
var client MQTT.Client

func BrokerConnection(id, username, password string) error {
	opts := MQTT.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", BrokerHost, Port))
	opts.SetClientID(id)
	opts.SetUsername(username)
	opts.SetPassword(password)

	opts.SetDefaultPublishHandler(messagePubHandler)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler

	client = MQTT.NewClient(opts)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		return token.Error()
	}

	return nil
}

var messagePubHandler MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
	log.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
}

var connectHandler MQTT.OnConnectHandler = func(client MQTT.Client) {
	log.Printf("Connected to broker: tcp://%s:%d\n", BrokerHost, Port)
}

var connectLostHandler MQTT.ConnectionLostHandler = func(client MQTT.Client, err error) {
	log.Printf("Broker connection lost: %v", err)
}

func Publish(t string, m string) {
	token := client.Publish(t, 0, false, m)
	token.Wait()
}

func Sub(t string) {
	token := client.Subscribe(t, 1, nil)
	token.Wait()
	log.Printf("Subscribed to topic: %s\n", t)
}
