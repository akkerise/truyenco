package kafka

import "github.com/segmentio/kafka-go"

func NewWriter(addr, topic string) *kafka.Writer {
    return &kafka.Writer{
        Addr:     kafka.TCP(addr),
        Topic:    topic,
        Balancer: &kafka.LeastBytes{},
    }
}

func NewReader(addr, topic, groupID string) *kafka.Reader {
    return kafka.NewReader(kafka.ReaderConfig{
        Brokers: []string{addr},
        Topic:   topic,
        GroupID: groupID,
    })
}

