package queue

import (
	"bs.mobgi.cc/app/vars"
	"github.com/Shopify/sarama"
)

type KafkaProducer struct {
	c     sarama.SyncProducer
	topic string
	key   string
}

func NewKafkaProducer(topic, key string) (*KafkaProducer, error) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll        // ack应答机制
	config.Producer.Partitioner = sarama.NewHashPartitioner // 发送分区
	config.Producer.Return.Successes = true                 // 回复确认
	// 连接kafka
	client, err := sarama.NewSyncProducer([]string{vars.YmlConfig.GetString("Kafka.Host")}, config)
	if err != nil {
		return nil, err
	}

	return &KafkaProducer{c: client, topic: topic, key: key}, nil
}

func (p *KafkaProducer) SendMessages(setMsg func(string, sarama.StringEncoder) []*sarama.ProducerMessage) (err error) {
	defer p.c.Close()

	// 发送消息
	err = p.c.SendMessages(setMsg(p.topic, sarama.StringEncoder(p.key)))
	return
}

func (p *KafkaProducer) SendMessage(msg []byte) (err error) {
	defer p.c.Close()

	// 发送消息
	_, _, err = p.c.SendMessage(&sarama.ProducerMessage{
		Topic: p.topic,
		Key:   sarama.StringEncoder(p.key),
		Value: sarama.ByteEncoder(msg),
	})
	return
}
