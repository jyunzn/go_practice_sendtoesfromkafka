package config


type LogTransferConf struct {
	Kafka `ini:"kafka"`
	ES `ini:"es"`
}

type Kafka struct {
	Address string `ini:"address"`
	Topic string `ini:"topic"`
}

type ES struct {
	Address string `ini:"address"`
	GrNumbers uint `ini:"gr_numbers"`
	BufSize uint   `ini:"buf_size"`
}
