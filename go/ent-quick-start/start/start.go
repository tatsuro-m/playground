package main

import (
	"context"
	"entqs/ent"
	"entqs/ent/car"
	"entqs/ent/group"
	"entqs/ent/migrate"
	"entqs/ent/user"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	client, err := ent.Open("mysql", getDSN())
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()

	ctx := context.Background()
	client.Schema.WriteTo(ctx, os.Stdout, migrate.WithDropColumn(true), migrate.WithDropIndex(true))

	err = client.Schema.Create(
		ctx, migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
		migrate.WithGlobalUniqueID(true))
	if err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	//CreateUser(ctx, client)
	//users, _ := QueryUser(ctx, client)
	//CreateCars(ctx, client)
	//QueryCars(ctx, users[len(users)-1])
	//QueryCarUsers(ctx, users[len(users)-1])

	//CreateGraph(ctx, client)
	//QueryGithub(ctx, client)
	//QueryArielCars(ctx, client)
	//QueryGroupWithUsers(ctx, client)

	CreateTest(ctx, client)
}

func CreateTest(ctx context.Context, client *ent.Client) error {
	t, err := client.Test.Create().SetTest(true).Save(ctx)
	if err != nil {
		return err
	}

	fmt.Println("created test: ", t)
	return nil
}

func QueryGroupWithUsers(ctx context.Context, client *ent.Client) error {
	groups, err := client.Group.
		Query().
		Where(group.HasUsers()).
		All(ctx)
	if err != nil {
		return fmt.Errorf("failed getting groups: %w", err)
	}
	log.Println("groups returned:", groups)
	// Output: (Group(Name=GitHub), Group(Name=GitLab),)
	return nil
}

func QueryArielCars(ctx context.Context, client *ent.Client) error {
	// 前の手順でユーザー "Ariel" を入手する
	a8m := client.User.
		Query().
		Where(
			user.HasCars(),
			user.Name("Ariel"),
		).
		OnlyX(ctx)
	cars, err := a8m. // a8mが接続されているグループを取得する
				QueryGroups(). // (Group(Name=GitHub), Group(Name=GitLab),)
				QueryUsers().  // (User(Name=Ariel, Age=30), User(Name=Neta, Age=28),)
				QueryCars().   //
				Where(         //
			car.Not( //  NetaとArielの車を取得する
				car.Model("Mazda"), //  しかし、"Mazda"というモデル名の車は除外する
			), //
		). //
		All(ctx)
	if err != nil {
		return fmt.Errorf("failed getting cars: %w", err)
	}
	log.Println("cars returned:", cars)
	// Output: (Car(Model=Tesla, RegisteredAt=<Time>), Car(Model=Ford, RegisteredAt=<Time>),)
	return nil
}

func QueryGithub(ctx context.Context, client *ent.Client) error {
	cars, err := client.Group.
		Query().
		Where(group.Name("GitHub")). // (Group(Name=GitHub),)
		QueryUsers().                // (User(Name=Ariel, Age=30),)
		QueryCars().                 // (Car(Model=Tesla, RegisteredAt=<Time>), Car(Model=Mazda, RegisteredAt=<Time>),)
		All(ctx)
	if err != nil {
		return fmt.Errorf("failed getting cars: %w", err)
	}
	log.Println("cars returned:", cars)
	// Output: (Car(Model=Tesla, RegisteredAt=<Time>), Car(Model=Mazda, RegisteredAt=<Time>),)
	return nil
}

func CreateGraph(ctx context.Context, client *ent.Client) error {
	// 最初に、ユーザーを複数作成する
	a8m, err := client.User.
		Create().
		SetAge(30).
		SetName("Ariel").
		Save(ctx)
	if err != nil {
		return err
	}
	neta, err := client.User.
		Create().
		SetAge(28).
		SetName("Neta").
		Save(ctx)
	if err != nil {
		return err
	}
	// その後、車を複数作成し、作成中のユーザーと紐付けます
	err = client.Car.
		Create().
		SetModel("Tesla").
		SetRegisteredAt(time.Now()). // グラフ内の時間を無視する
		SetOwner(a8m).               // このグラフをユーザー"Ariel"にアタッチする
		Exec(ctx)
	if err != nil {
		return err
	}
	err = client.Car.
		Create().
		SetModel("Mazda").
		SetRegisteredAt(time.Now()). // グラフ内の時間を無視する
		SetOwner(a8m).               // このグラフをユーザー"Ariel"にアタッチする
		Exec(ctx)
	if err != nil {
		return err
	}
	err = client.Car.
		Create().
		SetModel("Ford").
		SetRegisteredAt(time.Now()). // グラフ内の時間を無視する
		SetOwner(neta).              // このグラフをユーザー"Neta"にアタッチする
		Exec(ctx)
	if err != nil {
		return err
	}
	// グループを作って、作成中にユーザーを追加します。
	err = client.Group.
		Create().
		SetName("GitLab").
		AddUsers(neta, a8m).
		Exec(ctx)
	if err != nil {
		return err
	}
	err = client.Group.
		Create().
		SetName("GitHub").
		AddUsers(a8m).
		Exec(ctx)
	if err != nil {
		return err
	}
	log.Println("The graph was created successfully")
	return nil
}

func CreateUser(ctx context.Context, client *ent.Client) (*ent.User, error) {
	u, err := client.User.
		Create().
		SetAge(30).
		SetName("a8m").
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating user: %w", err)
	}
	log.Println("user was created: ", u)
	return u, nil
}

func QueryUser(ctx context.Context, client *ent.Client) ([]*ent.User, error) {
	u, err := client.User.
		Query().
		Where(user.Name("a8m")).All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying user: %w", err)
	}
	fmt.Println("user returned: ", u)
	return u, nil
}

func CreateCars(ctx context.Context, client *ent.Client) (*ent.User, error) {
	// "Tesla"というモデルの車を新しく作成します
	tesla, err := client.Car.
		Create().
		SetModel("Tesla").
		SetRegisteredAt(time.Now()).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating car: %w", err)
	}
	log.Println("car was created: ", tesla)

	// "Ford"というモデルの車を新しく作成します
	ford, err := client.Car.
		Create().
		SetModel("Ford").
		SetRegisteredAt(time.Now()).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating car: %w", err)
	}
	log.Println("car was created: ", ford)

	// 新しいユーザーを作成し、2台の車を所有させます
	a8m, err := client.User.
		Create().
		SetAge(30).
		SetName("a8m").
		AddCars(tesla, ford).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating user: %w", err)
	}
	log.Println("user was created: ", a8m)
	return a8m, nil
}

func QueryCars(ctx context.Context, a8m *ent.User) error {
	cars, err := a8m.QueryCars().All(ctx)
	if err != nil {
		return fmt.Errorf("failed querying user cars: %w", err)
	}
	log.Println("returned cars:", cars)

	// 特定の車をフィルタリングするには
	ford, err := a8m.QueryCars().
		Where(car.Model("Ford")).
		Only(ctx)
	if err != nil {
		return fmt.Errorf("failed querying user cars: %w", err)
	}
	log.Println(ford)
	return nil
}

func QueryCarUsers(ctx context.Context, a8m *ent.User) error {
	cars, err := a8m.QueryCars().All(ctx)
	if err != nil {
		return fmt.Errorf("failed querying user cars: %w", err)
	}
	// 逆エッジを取得する
	for _, ca := range cars {
		owner, err := ca.QueryOwner().Only(ctx)
		if err != nil {
			return fmt.Errorf("failed querying car %q owner: %w", ca.Model, err)
		}
		log.Printf("car %q owner: %q\n", ca.Model, owner.Name)
	}
	return nil
}

func getDSN() string {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	return fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?parseTime=true",
		dbUser, dbPassword, os.Getenv("DB_HOST"), dbName)
}
