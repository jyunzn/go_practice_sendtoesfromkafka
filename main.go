package main


import (
	"fmt"

	"logtransfer/config"
	"logtransfer/kafka"
	"logtransfer/es"

	"gopkg.in/ini.v1"
)


func main() {

	// 加載配置
	var cnf config.LogTransferConf
	err := ini.MapTo(&cnf, "./config/config.ini")
	if err != nil {
		fmt.Println("加載配置失敗", err)
		return
	}
	fmt.Printf("%#v\n", cnf)


	// 初始化 es
	err = es.Init(cnf.ES.Address, cnf.ES.BufSize, cnf.ES.GrNumbers)
	if err != nil {
		fmt.Println("es init err: ", err)
		return
	}

	// 初始化 kafka
	consumer, err := kafka.Init([]string { cnf.Kafka.Address }, cnf.Kafka.Topic)
	if err != nil {
		fmt.Println("kafka init err: ", err)
		return
	}
	defer consumer.Close()
	select {}
}

