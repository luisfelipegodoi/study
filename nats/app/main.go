package main

import "nats/config"

func main() {
	var envs = config.LoadEnvironments()
	var services = config.InitServices(envs)
}
