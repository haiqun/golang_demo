package conf

// 配置项包

type AppConf struct {
	KafkaConfig   `ini:"kafka"`
	TaillogConfig `ini:"taillog"`
}

type KafkaConfig struct {
	Address string `ini:"address"`
	Topic   string `ini:"topic"`
}

type TaillogConfig struct {
	Filename string `ini:"filename"`
}
