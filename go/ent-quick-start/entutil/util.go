package entutil

import (
	"entqs/ent"
)

var c *ent.Client

func InitClient(client *ent.Client) *ent.Client {
	c = client
	return c
}

func GetEntClient() *ent.Client {
	return c
}
