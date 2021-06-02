package main

import (
	"fmt"
	"generator/config"
	mqtt "generator/mqtt"
	"math/rand"
	"time"
)

func init() {

	if err := config.LoadConfigFile(); err != nil {
		panic(err)
	}

	// setup mqtt broker and connect
	mqtt.BrokerHost = config.BrokerHost
	mqtt.Port = config.BrokerPort

	if err := mqtt.BrokerConnection(config.BrokerName, config.BrokerUser, config.BrokerPwd); err != nil {
		fmt.Printf("Broker error: %v", err)
	}

}

func influxData() (string, string) {

	rand.Seed(time.Now().UnixNano())

	ts := time.Now().UnixNano()

	x := rand.Int()
	y := rand.Int31()
	z := rand.Int31()
	t := rand.Float32()

	return fmt.Sprintf("vibration,type=velocity,area=JZ1,equipment=lm1,location=main_drive,component=front_bearing x=%v,y=%v,z=%v %v", x, y, z, ts),
		fmt.Sprintf("temperature,type=deg_c,area=JZ1,equipment=lm1,location=main_drive,component=front_bearing value=%v %v", t, ts)
}

func main() {
	var msg1, msg2 string

	for {
		msg1, msg2 = influxData()

		mqtt.Publish("mehm/jz1/lm1/motor1/vibration", msg1)
		mqtt.Publish("mehm/jz1/lm1/motor1/temperature", msg2)
		time.Sleep(time.Duration(config.Delay) * time.Millisecond)
	}
}
