package main

import (
	"context"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/sirupsen/logrus"
	"sync"
)


func kafka(	ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	// 根据给定的代理地址和配置创建一个消费者
	consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, nil)
	if err != nil {
		panic(err)
	}
	wg1 := sync.WaitGroup{}
	//Partitions(topic):该方法返回了该topic的所有分区id
	topics := []string{"test","producer"}
	//var mu sync.Mutex
	for _, value := range topics {
		partitionList, err := consumer.Partitions(value)
		if err != err {
			logrus.Error("get consumer partitions %s err:", value ,err)
		}
		for _, partition := range partitionList {
			wg1.Add(1)
			go func(value string, partition int32) {
				//ConsumePartition方法根据主题，分区和给定的偏移量创建创建了相应的分区消费者
				//如果该分区消费者已经消费了该信息将会返回error
				//sarama.OffsetNewest:表明了为最新消息
				pc, err := consumer.ConsumePartition( value, partition, sarama.OffsetNewest)
				if err != err {
					logrus.Error("get consumer %s err:", value ,err)
				}
				defer pc.AsyncClose()
				//Messages()该方法返回一个消费消息类型的只读通道，由代理产生
				for msg := range pc.Messages() {
					fmt.Printf("%s---Partition:%d, Offset:%d, Key:%s, Value:%s\n", msg.Topic, msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
				}
				wg1.Done()
			}(value, partition)
		}
	}
	<- ctx.Done()
}
func main(){
	wg := sync.WaitGroup{}
	wg.Add(1)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go kafka( ctx, &wg )
	wg.Wait()
	logrus.Error("===================")
}
