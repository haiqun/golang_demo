package conf

// 配置项包

type AppConf struct {
	KafkaConfig   `ini:"kafka"`
	//TaillogConfig `ini:"taillog"`
	EtcdConfig `ini:"etcd"`
}

type KafkaConfig struct {
	Address string `ini:"address"`
	//Topic   string `ini:"topic"`
}

type EtcdConfig struct {
	Endpoints string `ini:"endpoints"`
	Key string `ini:"conllect_log_key"`
	Timeout int `int:"timeout"`
}


type TaillogConfig struct {
	Filename string `ini:"filename"`
}
