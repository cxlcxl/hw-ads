package test

import (
	"bs.mobgi.cc/library/queue"
	"github.com/Shopify/sarama"
	"strconv"
	"testing"
)

func TestProducer(t *testing.T) {
	producer, err := queue.NewKafkaProducer("cronErrors", "country")
	if err != nil {
		t.Fatal(err)
	}
	a := []int{123, 456, 789}
	err = producer.SendMessages(func(topic string, key sarama.StringEncoder) (msg []*sarama.ProducerMessage) {
		for _, i := range a {
			msg = append(msg, &sarama.ProducerMessage{
				Topic: topic,
				Key:   key,
				Value: sarama.StringEncoder(strconv.Itoa(i)),
			})
		}
		return
	})
	t.Log(err)
}
