package user

import (
	"context"
	"entqs/ent"
	"entqs/entutil"
	"fmt"
	"log"
)

func Create(u *ent.User, ctx context.Context) (*ent.User, error) {
	user, err := entutil.GetEntClient().User.Create().SetAge(u.Age).SetName(u.Name).Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating user: %v", err)
	}
	log.Println("user was created: ", user)

	return user, nil
}
