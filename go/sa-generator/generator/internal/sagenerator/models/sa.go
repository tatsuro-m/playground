package models

import (
	"generator/internal/sagenerator/config"
)

type Sa struct {
	Env         string
	ServiceName string
}

func GetMatrix(c *config.Config) []*Sa {
	sas := make([]*Sa, 0)
	for _, env := range c.TargetEnvs {
		for _, service := range c.TargetServices {
			sas = append(sas, &Sa{Env: env, ServiceName: service})
		}
	}

	return sas
}
