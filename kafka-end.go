package abango

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Shopify/sarama"
	e "github.com/dabory/abango/etc"
)

//////////// Kafka EndPoint /////////////
func KafkaRequest(v *AbangoAsk) (string, string, error) {

	topic := XConfig["KafkaTopic"]

	conn := XConfig["KafkaAddr"] + ":" + XConfig["KafkaPort"]
	e.MyLog(v.HomeRoot+"abango.log", "C-A")
	svars := make(map[string]string)
	for _, p := range v.ServerParams {
		svars[p.Key] = p.Value
	}
	apiMethod := strings.ToUpper(svars["api_method"])
	e.MyLog(v.HomeRoot+"abango.log", "C-B")
	ReturnTopic := v.UniqueId

	askstr, _ := json.Marshal(&v)

	if _, _, err := KafkaProducer(string(askstr), topic, conn, apiMethod); err == nil {
		TmpInt, _ := strconv.Atoi(XConfig["KafkaCosumerTimeout"])
		timeout := int64(TmpInt)
		e.MyLog(v.HomeRoot+"abango.log", "C-C")
		if msg, err := KafkaReturnConsumer(ReturnTopic, conn, timeout); err == nil {
			e.MyLog(v.HomeRoot+"abango.log", "C-D")
			return msg, "202", nil // Normal Retrun
		} else {
			return "", "503", e.MyErr("ADFARQ#FA- Kafka Service Unavailable", err, true)
		}
	} else {
		return "", "503", e.MyErr("QWER!#$%^&*#- Kafka Server or Network disconnected", err, true) //
	}
}

func KafkaReturnConsumer(topic string, conn string, timeout int64) (string, error) {

	kfcf := sarama.NewConfig()
	kfcf.Consumer.Return.Errors = true

	// Specify brokers address. This is default one
	brokers := []string{conn}
	master, err := sarama.NewConsumer(brokers, kfcf)
	if err != nil {
		return "", e.MyErr("DSFYODGBWSVE-Kafka Consumer Not created", err, true)
	}

	defer func() {
		if err := master.Close(); err != nil {
			return
		}
	}()

	// How to decide partition, is it fixed value...?
	// consumer, err := master.ConsumePartition(retTopic, 0, sarama.OffsetNewest)
	// 다른 놈이 읽어가고 있는 것 같다. 그래서 항상 처음 것으로 하였다.
	consumer, err := master.ConsumePartition(topic, 0, sarama.OffsetOldest)
	if err != nil {
		e.MyErr("ConsumePartition-QERWRVZDFVNMRE", err, true)
	}

	strMsg := ""

	doneCh := make(chan struct{})
	ctx, cancel := context.WithCancel(context.Background())

	go func(ctx context.Context) {
		for {
			select {
			case err := <-consumer.Errors():
				fmt.Println(err)
			case msg := <-consumer.Messages():
				strMsg = string(msg.Value)
				doneCh <- struct{}{}
			case <-ctx.Done(): // if cancel() execute
				doneCh <- struct{}{}
				return
			}
		}
	}(ctx)

	go func() {
		time.Sleep(time.Duration(timeout) * time.Second)
		e.OkLog("Timeout-Kafka connection after " + strconv.Itoa(int(timeout)) + " Seconds")
		cancel()
	}()

	<-doneCh

	return strMsg, nil
}
