package abango

import (
	"encoding/json"
	"fmt"

	"github.com/Shopify/sarama"
	e "github.com/dabory/abango/etc"
)

//////////// Kafka Service /////////////
func KafkaSvcStandBy(RouterHandler func(*AbangoAsk)) {

	kfcf := sarama.NewConfig()
	kfcf.Consumer.Return.Errors = true

	conn := XConfig["KafkaAddr"] + ":" + XConfig["KafkaPort"]
	brokers := []string{conn}

	// Create new consumer
	master, err := sarama.NewConsumer(brokers, kfcf)
	if err != nil {
		e.MyErr("QRVAQEAREADVSQ-Kafka Consumer Not created", err, true)
		return
	}

	defer func() {
		if err := master.Close(); err != nil {
			e.MyErr("IRJDNWRTSE-Kafka Consumer Not closed", err, true)
			return
		}
	}()

	// How to decide partition, is it fixed value...?
	topic := XConfig["KafkaTopic"]
	if consumer, err := master.ConsumePartition(topic, 0, sarama.OffsetNewest); err == nil {
		// e.AokLog("Kafka:[topic->" + topic + "] StandBy !")
		doneCh := make(chan struct{})
		go func() {
			for {
				select {
				case err := <-consumer.Errors():
					fmt.Println(err)
				case msg := <-consumer.Messages():
					var v AbangoAsk
					e.Tp(msg.Value)
					e.Tp(string(msg.Value))
					if err := json.Unmarshal(msg.Value, &v); err == nil {
						e.OkLog("Kafka API [" + v.AskName + "]-[" + topic + "]-[" + v.UniqueId + "]")
						RouterHandler(&v)
					} else {
						e.MyErr("QEWRQCVZVCXER-Unmarshall-msg.Value", err, false)
					}
				}
			}
		}()
		<-doneCh

	} else {
		e.MyErr("ConsumePartition-WERDCAYQDFARW", err, true)
	}
}
