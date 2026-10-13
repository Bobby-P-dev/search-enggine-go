package mqttclient

import (
	"errors"
	"fmt"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

const (
	DefaultBrokerURL = "tcp://test.mosquitto.org:1883"
	DefaultTopic     = "smarthome/lab-a7f48e/lamp-bedroom/set"
)

type Publisher struct {
	client mqtt.Client
	topic  string
}

func NewPublisher(
	brokerURL string,
	topic string,
) (*Publisher, error) {
	if brokerURL == "" {
		brokerURL = DefaultBrokerURL
	}

	if topic == "" {
		topic = DefaultTopic
	}

	options := mqtt.NewClientOptions().
		AddBroker(brokerURL).
		SetClientID(
			fmt.Sprintf("smart-home-go-%d", time.Now().UnixNano()),
		).
		SetAutoReconnect(true)

	client := mqtt.NewClient(options)
	token := client.Connect()

	if !token.WaitTimeout(10 * time.Second) {
		return nil, errors.New("MQTT connection timeout")
	}

	if err := token.Error(); err != nil {
		return nil, fmt.Errorf("MQTT connect: %w", err)
	}

	return &Publisher{
		client: client,
		topic:  topic,
	}, nil
}

func (p *Publisher) PublishLampCommand(state string) error {
	if state != "ON" && state != "OFF" {
		return errors.New("invalid lamp state")
	}

	if !p.client.IsConnected() {
		return errors.New("MQTT broker disconnected")
	}

	token := p.client.Publish(
		p.topic,
		1,
		false,
		state,
	)

	if !token.WaitTimeout(5 * time.Second) {
		return errors.New("MQTT publish timeout")
	}

	return token.Error()
}

func (p *Publisher) Close() {
	p.client.Disconnect(250)
}
