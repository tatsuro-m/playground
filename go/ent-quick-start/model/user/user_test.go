package user_test

import (
	"context"
	"entqs/ent"
	"entqs/model/user"
	"entqs/thelper"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func TestCreate(t *testing.T) {
	table := []struct {
		name    string
		user    *ent.User
		want    *ent.User
		wantErr bool
	}{
		{
			name:    "正常に作成されること",
			user:    &ent.User{Age: 22, Name: "test1"},
			want:    &ent.User{Age: 22, Name: "test1"},
			wantErr: false,
		},
	}

	for _, dd := range table {
		t.Run(dd.name, func(t *testing.T) {
			client := thelper.InitEntClient(t)
			defer client.Close()

			got, err := user.Create(dd.user, context.Background())
			log.Println(got)

			if dd.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			if d := cmp.Diff(got, dd.want, cmpopts.IgnoreFields(*got, "config", "ID")); len(d) != 0 {
				t.Errorf("test failed want:　%v got: %v\ndiff: %s", dd.want, got, d)
			}
		})
	}
}
