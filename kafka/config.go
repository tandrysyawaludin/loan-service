package kafka

type Config struct {
	Brokers []string
	GroupID string
	Topics  []string
}

func DefaultConfig() *Config {
	return &Config{
		Brokers: []string{"localhost:9092"},
		GroupID: "loan-service-group",
		Topics:  []string{"loan-events"},
	}
}
