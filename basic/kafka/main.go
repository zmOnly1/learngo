package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"strings"
	"time"
)

const topic = "nginx_log"

func main() {
	produce()
	consumer()
}

func consumer() {
	consumer, err := sarama.NewConsumer(strings.Split("localhost:9092", ","), nil)
	if err != nil {
		panic(err)
	}
	partitionList, err := consumer.Partitions(topic)
	if err != nil {
		panic(err)
	}
	fmt.Println(partitionList)
	for partition := range partitionList {
		pc, err := consumer.ConsumePartition(topic, int32(partition), sarama.OffsetNewest)
		if err != nil {
			panic(err)
		}
		defer pc.AsyncClose()
		go func(pc sarama.PartitionConsumer) {
			for msg := range pc.Messages() {
				fmt.Printf("Partition:%d, Offset:%d, key:%s, Value:%s\n", msg.Partition, msg.Offset,
					string(msg.Key), string(msg.Value))
			}
		}(pc)
	}
	fmt.Println("done")
	time.Sleep(time.Hour)
	consumer.Close()
}

func produce() {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true

	msg := &sarama.ProducerMessage{}
	msg.Topic = topic
	msg.Value = sarama.StringEncoder("this is a good test, my message is good")

	client, err := sarama.NewSyncProducer([]string{"localhost:9092"}, config)
	if err != nil {
		panic(err)
	}
	defer client.Close()
	pid, offset, err := client.SendMessage(msg)
	if err != nil {
		panic(err)
	}
	fmt.Printf("pid:%v offset:%v\n", pid, offset)
}
