package config

type TopicConfig struct {
	IntradayValues string `envconfig:"KAFKA_INTRADAY_VALUE" default:"gpsi.stocks.intraday.value"`
}

type SchemaRegistryConfig struct {
	URL                 string `envconfig:"KAFKA_REGISTRY_URL" default:"http://localhost:8081"`
	Username            string `envconfig:"KAFKA_REGISTRY_USER"`
	Password            string `envconfig:"KAFKA_REGISTRY_PASSWORD"`
	AutoRegisterSchemas bool   `envconfig:"KAFKA_REGISTRY_AUTO_REGISTER" default:"false"`
}

type KafkaConfig[T any] struct {
	Brokers                []string `envconfig:"KAFKA_BROKER_HOSTS" default:"localhost:9092"`
	GroupID                string   `envconfig:"KAFKA_CONSUMER_GROUP_ID"`
	Username               string   `envconfig:"KAFKA_USER"`
	Password               string   `envconfig:"KAFKA_PASSWORD"`
	AllowAutoTopicCreation bool     `envconfig:"KAFKA_AUTO_TOPIC_CREATION" default:"false"`
	Topics                 T
	SchemaRegistry         SchemaRegistryConfig
}
