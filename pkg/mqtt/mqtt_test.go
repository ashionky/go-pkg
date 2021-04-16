/**
 * @Author pibing
 * @create 2021/3/24 10:40 AM
 */

package mqtt

import (
	"encoding/json"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"sync"
	"testing"
	"time"
)

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Connected")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Connect lost: %v", err)
}

func TestMqtt(t *testing.T) {
	var broker = "broker.emqx.io"
	var port = 1883
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", broker, port))
	opts.SetClientID("go_mqtt_client")
	opts.SetUsername("emqx")
	opts.SetPassword("public")
	opts.SetDefaultPublishHandler(messagePubHandler)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	sub(client)
	publish(client)

	client.Disconnect(250)
}

func publish(client mqtt.Client) {
	num := 10
	for i := 0; i < num; i++ {
		text := fmt.Sprintf("Message %d", i)
		token := client.Publish("topic/test", 0, false, text)
		token.Wait()
		time.Sleep(time.Second)
	}
}

func sub(client mqtt.Client) {
	topic := "topic/test"
	token := client.Subscribe(topic, 1, nil)
	token.Wait()
	fmt.Printf("Subscribed to topic: %s \n", topic)
}

func TestNewClient(t *testing.T) {
	var (
		clientId = "test_clientid"
		wg       sync.WaitGroup
	)
	client := NewClient(clientId)
	err := client.Connect()
	if err != nil {
		t.Errorf(err.Error())
	}

	wg.Add(1)
	go func() {
		err := client.Subscribe(func(c *Client, msg *Message) {
			fmt.Printf("接收到消息: %+v \n", msg)
			wg.Done()
		}, 1, "mqtt")

		if err != nil {
			panic(err)
		}
	}()

	msg := &Message{
		ClientID: clientId,
		Type:     "text",
		Data:     "Hello test_client",
		Time:     time.Now().Unix(),
	}
	data, _ := json.Marshal(msg)

	err = client.Publish("mqtt", 1, false, data)
	if err != nil {
		panic(err)
	}

	wg.Wait()

}
