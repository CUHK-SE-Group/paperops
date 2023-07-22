package main

import (
	"context"
	"encoding/json"
	"time"

	"github.com/twmb/franz-go/pkg/kgo"
)

type Producer struct {
	client *kgo.Client
	topic  string
}

func NewProducer(brokers []string, topic string) *Producer {
	client, err := kgo.NewClient(
		kgo.SeedBrokers(brokers...),
	)
	if err != nil {
		panic(err)
	}
	return &Producer{client: client, topic: topic}
}
func (p *Producer) SendMessage(msg Entry) error {
	ctx := context.Background()
	b, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	p.client.Produce(ctx, &kgo.Record{Topic: p.topic, Value: b}, nil)
	return nil
}
func (p *Producer) Flush() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	p.client.Flush(ctx)
}
func (p *Producer) Close() {

	p.client.Close()
}
