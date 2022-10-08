package es

import (
	"context"
	"fmt"
	"time"
	"github.com/olivere/elastic/v7"
)

type LogData struct {
	Data string `json:"data"`
}

type LogDataContainer struct {
	Data LogData
	Topic string
	TypeStr string
}

var (
	client *elastic.Client
	buffer chan *LogDataContainer
)

func Init(address string, bufSize uint, grNumbers uint) (err error) {
	client, err = elastic.NewClient(elastic.SetURL("http://" + address))
	if err != nil {
		return
	}


	buffer = make(chan *LogDataContainer, bufSize)
	for i := 0; i <= int(grNumbers); i++ {
		go SendToES()
	}
	return
}



func SendToESBuffer(ldc *LogDataContainer) {
	buffer <- ldc
}

func SendToES() {
	for {
		select {
			case ldc := <- buffer:
				_, err := client.Index().
					Index(ldc.Topic).
					Type(ldc.TypeStr).
					BodyJson(ldc.Data).
					Do(context.Background())

				if err != nil {
					fmt.Println("Send Log To Es ERROR: ", err)
					continue
				}
			default:
				time.Sleep(time.Second)
		}
	}
}
