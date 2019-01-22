package config

import "github.com/spf13/viper"

const (
	configType = "yml"
)

// Configuration is structure where application properties are stored.
type Configuration struct {
	GrpcClient *GrpcServer `mapstructure:"grpc_server"`
}

// GrpcServer configuration.
type GrpcServer struct {
	Address string
}

// Parse retrieves configuration from configuration file and maps it to Configuration.
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
