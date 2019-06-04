package apploader

import (
	"github.com/vitaliyyevenko/continuum-utils/RMM/preparation/generation"
	"github.com/vitaliyyevenko/continuum-utils/RMM/preparation/persistency/kafka"
	"github.com/vitaliyyevenko/continuum-utils/RMM/preparation/persistency/postgres"
	"github.com/vitaliyyevenko/continuum-utils/RMM/preparation/rest"
	"github.com/vitaliyyevenko/continuum-utils/RMM/preparation/writer"
)

var restClients = []func(rest.Environment) error{
	rest.LoadAsset,
	rest.LoadEntitlement,
}

var clients = []func() error{
	generation.Load,
	writer.Load,
	postgres.Load,
	kafka.Load,
}

// Load initializes clients
func Load(env rest.Environment) (err error) {
	for _, rc := range restClients {
		err = rc(env)
		if err != nil {
			return err
		}
	}
	for _, rc := range clients {
		err = rc()
		if err != nil {
			return err
		}
	}

	return nil
}
