package module

import (
	"fmt"
	"github.com/Corner-W/sk/log"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"time"
)

var (
	MqClient mqtt.Client
)

type MqttTask struct {
	msg string
}

func NewMqtt() *MqttTask {

	return &MqttTask{
		msg: "mqtt module",
	}
}

func (a *MqttTask) OnInit() {

	log.Debug("module mqtt init...")

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
	MqClient := mqtt.NewClient(opts)
	if token := MqClient.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

}

func (a *MqttTask) OnDestroy() {

}

func (a *MqttTask) MsgProc(closeSig chan bool, message interface{}) {

	log.Debug("module mqtt  Enter...")

}

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Connected remote addr: ")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Connect lost: %v", err)
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
	fmt.Printf("Subscribed to topic: %s", topic)
}
