package config

import "github.com/nats-io/nats.go"

type Services struct{
	NatsJetStream nats.JetStreamContext
}

func InitServices(envs *Environments) *Services{

	return &Services{}
}

func initNatsJetStream(envs *Environments) nats.JetStreamContext{
	
}