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
		{
			name:    "age を指定しないと作成できないこと",
			user:    &ent.User{Name: "test1"},
			want:    &ent.User{Name: "test1"},
			wantErr: true,
		},
		{
			name:    "age は0だと作成できないこと",
			user:    &ent.User{Age: 0, Name: "test1"},
			want:    &ent.User{Age: 0, Name: "test1"},
			wantErr: true,
		},
		{
			name:    "name は空文字でも良いこと",
			user:    &ent.User{Age: 22},
			want:    &ent.User{Age: 22, Name: ""},
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
				if d := cmp.Diff(got, dd.want, cmpopts.IgnoreFields(*got, "config", "Edges", "ID")); len(d) != 0 {
					t.Errorf("test failed want:　%v got: %v\ndiff: %s", dd.want, got, d)
				}
			}
		})
	}
}

func TestUserToCarsRelation(t *testing.T) {
	t.Run("リレーションが正しく定義されていること", func(t *testing.T) {
		client := thelper.InitEntClient(t)
		defer client.Close()

		ctx := context.Background()
		tesla, err := client.Car.Create().SetModel("Tesla").Save(ctx)
		assert.NoError(t, err)
		toyota, err := client.Car.Create().SetModel("Toyota").Save(ctx)
		assert.NoError(t, err)

		u := &ent.User{Age: 22, Name: "test1"}
		user, err := user.Create(u, ctx)
		assert.NoError(t, err)

		user, err = client.User.UpdateOne(user).AddCars(tesla, toyota).Save(ctx)
		assert.NoError(t, err)

		cars, err := user.QueryCars().All(ctx)
		assert.NoError(t, err)
		assert.Equal(t, 2, len(cars))
	})
}
