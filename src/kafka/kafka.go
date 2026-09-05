package kafka

import (
	"context"
	"fmt"
	"learning-go/src/number_sequence"
	"sync"
	"time"

	"github.com/twmb/franz-go/pkg/kgo"
)

var ctx = context.Background()
var hosts = []string{"localhost:9092"}

const topic = "numeric-topic"

func produceRecord(client *kgo.Client, wg *sync.WaitGroup, number int) {
	// Wait a half second for each number
	time.Sleep(500 * time.Millisecond)

	record := &kgo.Record{Topic: topic, Value: fmt.Append(nil, number)}
	client.Produce(ctx, record, func(_ *kgo.Record, err error) {
		defer wg.Done()
		if err != nil {
			fmt.Printf("record had a produce error: %v\n", err)
		}
	})
}

func produceSequence(client *kgo.Client, wg *sync.WaitGroup, sequence *number_sequence.NumberSequence, numMsg int) {
	for range numMsg {
		produceRecord(client, wg, sequence.Next())
	}
}

func RunProducer(sequence *number_sequence.NumberSequence, numMsg int) {
	client, err := kgo.NewClient(kgo.SeedBrokers(hosts...))

	if err != nil {
		panic(err)
	}

	defer client.Close()

	var wg sync.WaitGroup
	wg.Add(numMsg)

	produceSequence(client, &wg, sequence, numMsg)
	wg.Wait()

	// Wait a second only for consumer consumes all messages
	time.Sleep(1 * time.Second)
}

func RunConsumer() {
	client, err := kgo.NewClient(kgo.SeedBrokers(hosts...), kgo.ConsumeTopics(topic))

	if err != nil {
		panic(err)
	}

	defer client.Close()

	for {
		fetches := client.PollFetches(ctx)
		if errs := fetches.Errors(); len(errs) > 0 {
			panic(fmt.Sprint(errs))
		}

		iter := fetches.RecordIter()
		for !iter.Done() {
			record := iter.Next()
			fmt.Println(string(record.Value))
		}
	}
}
