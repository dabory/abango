package abango

import (
	"github.com/Shopify/sarama"
	e "github.com/dabory/abango/etc"
)

func KafkaProducer(message string, topic string, conn string, apiMethod string) (int32, int64, error) {

	kfcf := sarama.NewConfig()
	kfcf.Producer.Retry.Max = 5
	kfcf.Producer.RequiredAcks = sarama.WaitForAll
	kfcf.Producer.Return.Successes = true

	if apiMethod == "" { // Default is POST
		apiMethod = "ASYNC"
	}

	if apiMethod == "ASYNC" {
		e.MyLog("/home/end_abango/abango.log", "here-1")
		if prd, err := sarama.NewAsyncProducer([]string{conn}, kfcf); err == nil {
			e.MyLog("/home/end_abango/abango.log", "here-2")

			prd.Input() <- &sarama.ProducerMessage{
				Topic: topic,
				Value: sarama.StringEncoder(message), //[]byte doesn't work.
			}

			return 0, 0, nil
		} else {
			return 0, 0, e.MyErr("QEJHDRTTRRW-Kafka-NewSyncProducer-End", err, true)
		}

	} else if apiMethod == "SYNC" {
		if prd, err := sarama.NewSyncProducer([]string{conn}, kfcf); err == nil {
			msg := &sarama.ProducerMessage{
				Topic: topic,
				Value: sarama.StringEncoder(message), //[]byte doesn't work.
			}
			if part, offset, err := prd.SendMessage(msg); err == nil {
				return part, offset, nil
			} else {
				return 0, 0, e.MyErr("QEJIOPRTRRTRRW-Kafka-Sync-SendMessage", err, true)
			}

		} else {
			return 0, 0, e.MyErr("QEJHGTRSDRTTRRW-Kafka-NewSyncProducer-End", err, true)
		}

	} else {
		return 0, 0, e.MyErr("QEJHGTRSW-Kafka-ApiMethod Not available-End", nil, true)
	}
}
