package conf

// 配置项包

type AppConf struct {
	KafkaConfig   `ini:"kafka"`
	EsConfig `ini:"es"`
}

type KafkaConfig struct {
	Address string `ini:"address"`
	Topic string `ini:"topic"`
}

type EsConfig struct {
	Address string `ini:"address"`
}

