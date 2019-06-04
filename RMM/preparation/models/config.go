package models

import "github.com/vitaliyyevenko/continuum-utils/RMM/preparation/rest"

// Configuration describes paramenters to launch MS
type Configuration struct {
	SetUpDB     bool
	SendToKafka bool
	Port        int
	Environment rest.Environment
}
