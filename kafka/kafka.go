package kafka

import (
	// "fmt"

	"github.com/Shopify/sarama"
	"logtransfer/es"
)


func Init(host []string, topic string) (consumer sarama.Consumer, err error) {
	consumer, err = sarama.NewConsumer(host, nil)
	if err != nil {
		return
	}

	//
	partitionIDs, err := consumer.Partitions(topic)
	if err != nil {
		return
	}

	for _, partitionID := range partitionIDs {
		var pc sarama.PartitionConsumer
		pc, err = consumer.ConsumePartition(topic, int32(partitionID), sarama.OffsetNewest)
		if err != nil {
			return
		}

		go func (pc sarama.PartitionConsumer) {
			defer pc.AsyncClose()
			for msg := range pc.Messages() {

				ld := &es.LogDataContainer {
					Data: es.LogData {
						Data: string(msg.Value),
					},
					Topic: topic,
					TypeStr: "godjj",
				}

				es.SendToESBuffer(ld)
			}
		}(pc)

	}

	return

}
