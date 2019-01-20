package config

import "github.com/spf13/viper"

const (
	configType = "yml"
)

type Configuration struct {
	GrpcClient *GrpcServer `mapstructure:"grpc_server"`
}

type GrpcServer struct {
	Address string
}

func Parse() (*Configuration, error) {
	manager := viper.New()
	manager.AutomaticEnv()

	manager.SetConfigType(configType)
	manager.AddConfigPath(".")

	if err := manager.ReadInConfig(); err != nil {
		return nil, err
	}

	configuration := new(Configuration)
	if err := mapToStructure(manager, configuration); err != nil {
		return nil, err
	}

	return configuration, nil
}

func mapToStructure(manager *viper.Viper, configuration *Configuration) error {
	return manager.Unmarshal(configuration)
}