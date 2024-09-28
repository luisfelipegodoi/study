package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Environments struct {
	EnvironmentMode    string
	NatsAddr           string
	NatsStreamName     string
	NatsStreamSubjects string
	NatsSubjectFanout  string
	NatsSubjectDLQ     string
}

func LoadEnvironments() *Environments {
	viper.AddConfigPath(".")
	viper.AddConfigPath("../")
	viper.AddConfigPath("../../")
	viper.AddConfigPath("../../..")
	viper.AddConfigPath("./app")
	viper.SetConfigName("env")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("error to load environments from file: %v", err.Error())
	}

	fmt.Printf("successfull loaded environments")

	return &Environments{
		EnvironmentMode:    viper.GetString("environment.mode"),
		NatsAddr:           viper.GetString("nats.addrs"),
		NatsStreamName:     viper.GetString("nats.stream.name"),
		NatsStreamSubjects: viper.GetString("nats.stream.subjects"),
		NatsSubjectFanout:  viper.GetString("nats.stream.subjectFanout"),
		NatsSubjectDLQ:     viper.GetString("nats.stream.subjectDLQ"),
	}
}
