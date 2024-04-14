package config

import (
	"fmt"
	"os"

	"github.com/go-yaml/yaml"
)

type YamlParameters struct {
	confPath   string
	parameters map[string]string
}

type deploymentYamlScheme struct {
	Contents map[string]map[string]string `yaml:"env"`
}

type Config struct {
	ListenerHost     string `env:"HTTP_PORT"`
	ListenerHttpPort string `env:"LISTENER_HTTP_PORT"`
	SQLConnectionUrl string `env:""`
	App              App
}

type App struct {
	Debug bool `env:"GIN_DEBUG"`
}

func LoadConfig() (Config, error) {
	source := os.Getenv("CONFIG_SOURCE")
	if source == "local" {
		configuration, err := getConfigFromLocal()
		if err != nil {
			return Config{}, err
		}

		return configuration, nil
	}
	return Config{}, nil
}

func getConfigFromLocal() (Config, error) {
	deployParameters, err := getConfigFromYaml()
	if err != nil {
		return Config{}, fmt.Errorf("failed to parse local deployment config: %w", err)
	}

	config := Config{}

	for _, parameters := range deployParameters {
		switch parameters.confPath {
		case "LISTENER_HTTP_PORT":
			config.ListenerHttpPort = parameters.parameters["value"]
		case "LISTENER_HTTP_HOST":
			config.ListenerHost = parameters.parameters["value"]
		case "SQL_CONNECTION_URL":
			config.SQLConnectionUrl = parameters.parameters["value"]
		}
	}

	return config, nil
}

func getConfigFromYaml() ([]YamlParameters, error) {
	path := os.Getenv("DEPLOY_CONF")
	if path == "" {
		return nil, fmt.Errorf("environment variable \"DEPLOY_CONF\" must contain path to the yaml deployment file")
	}

	deploymentConfigContent, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read file with deployment config: %w", err)
	}

	var deploymentConfig deploymentYamlScheme
	if err := yaml.Unmarshal(deploymentConfigContent, &deploymentConfig); err != nil {
		return nil, fmt.Errorf("failed to read file with deployment config: %w", err)
	}

	result := []YamlParameters{}
	for configPath, configParameters := range deploymentConfig.Contents {
		result = append(result, YamlParameters{
			confPath:   configPath,
			parameters: configParameters,
		})
	}

	return result, nil
}

func IsEmpty(c Config) bool {
	return c == Config{}
}
