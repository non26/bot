package config

import "github.com/spf13/viper"

type Config struct {
	Port           int      `yaml:"port" json:"port"`
	Environment    string   `yaml:"environment" json:"environment"`
	HealthCheckMsg string   `yaml:"healthCheckMsg" json:"healthCheckMsg"`
	BNBotBFF       BNBotBFF `yaml:"bnBotBff" json:"bnBotBff"`
	// BNBotCore      BNBotCore `yaml:"bnBotCore" json:"bnBotCore"`
	DynamoDB DynamoDB `yaml:"dynamodb" json:"dynamodb"`
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) IsLocal() bool {
	return c.Environment == "local"
}

type DynamoDB struct {
	Region   string `yaml:"region" json:"region"`
	Ak       string `yaml:"ak" json:"ak"`
	Sk       string `yaml:"sk" json:"sk"`
	Endpoint string `yaml:"endpoint" json:"endpoint"`
}

type BNBotBFF struct {
	BotTradeManagement BNBotBFFBotTradeManagement `yaml:"botTradeManagement" json:"botTradeManagement"`
	BotManagement      BNBotBFFBotManagement      `yaml:"botManagement" json:"botManagement"`
}

type BNBotBFFBotTradeManagement struct {
	BaseURL          string `yaml:"baseurl" json:"baseurl"`
	NewOrderEndpoint string `yaml:"newOrderEndpoint" json:"newOrderEndpoint"`
}

type BNBotBFFBotManagement struct {
	BaseURL        string `yaml:"baseurl" json:"baseurl"`
	GetEndpoint    string `yaml:"endpoint" json:"endpoint"`
	UpdateEndpoint string `yaml:"updateEndpoint" json:"updateEndpoint"`
}

// type BNBotCore struct {
// 	BotOpening BNBotCoreBotOpening `yaml:"botOpening" json:"botOpening"`
// }

// type BNBotCoreBotOpening struct {
// 	BaseURL        string `yaml:"baseurl" json:"baseurl"`
// 	GetEndpoint    string `yaml:"getEndpoint" json:"getEndpoint"`
// 	GetAllEndpoint string `yaml:"getAllEndpoint" json:"getAllEndpoint"`
// }

func LoadConfig() (*Config, error) {
	config := NewConfig()
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	viper.Unmarshal(&config)
	return config, nil
}
