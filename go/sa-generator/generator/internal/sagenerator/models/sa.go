package models

import (
	"generator/internal/sagenerator/config"
)

type Sa struct {
	Env         string
	ServiceName string
}

func GetMatrix(c *config.Config) *[]Sa {
	return nil
}
