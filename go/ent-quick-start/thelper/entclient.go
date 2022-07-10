package thelper

import (
	"entqs/ent"
	"entqs/ent/enttest"
	"entqs/entutil"
	"testing"
)

func InitEntClient(t *testing.T) *ent.Client {
	t.Helper()
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	entutil.InitClient(client)

	return client
}
