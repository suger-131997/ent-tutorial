package main

import (
	"context"
	"ent-tutorial/ent"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	client, err := ent.Open("mysql", "docker:docker@tcp(localhost:3306)/test_database?parseTime=True")
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}
	defer client.Close()

	ctx := context.Background()
	client.Pet.Delete().Exec(ctx)
	client.User.Delete().Exec(ctx)
	client.Group.Delete().Exec(ctx)

	g1 := CreateGroup(ctx, client, "group1")
	g2 := CreateGroup(ctx, client, "group2")
	g3 := CreateGroup(ctx, client, "group3")
	u1 := CreateUser(ctx, client, "user1", time.Now(), g1)
	u2 := CreateUser(ctx, client, "user2", time.Now(), g1)
	u3 := CreateUser(ctx, client, "user3", time.Now(), g1)
	u4 := CreateUser(ctx, client, "user4", time.Now(), g2)
	u5 := CreateUser(ctx, client, "user5", time.Now(), g2)
	u6 := CreateUser(ctx, client, "user6", time.Now(), g3)
	CreatePet(ctx, client, "pet1", 1, u1)
	CreatePet(ctx, client, "pet2", 2, u1)
	CreatePet(ctx, client, "pet3", 3, u1)
	CreatePet(ctx, client, "pet4", 4, u2)
	CreatePet(ctx, client, "pet5", 5, u2)
	CreatePet(ctx, client, "pet6", 6, u3)
	CreatePet(ctx, client, "pet7", 7, u4)
	CreatePet(ctx, client, "pet8", 8, u5)
	CreatePet(ctx, client, "pet9", 9, u6)
}

func CreateGroup(ctx context.Context, client *ent.Client, name string) *ent.Group {
	g, err := client.Group.
		Create().
		SetName(name).
		Save(ctx)
	if err != nil {
		panic(err)
	}
	return g
}

func CreateUser(ctx context.Context, client *ent.Client, name string, registeredAt time.Time, group *ent.Group) *ent.User {
	u, err := client.User.
		Create().
		SetName(name).
		SetRegisteredAt(registeredAt).
		AddGroups(group).
		Save(ctx)

	if err != nil {
		panic(err)
	}
	return u
}

func CreatePet(ctx context.Context, client *ent.Client, name string, age int, owner *ent.User) *ent.Pet {
	p, err := client.Pet.
		Create().
		SetName(name).
		SetAge(age).
		SetOwner(owner).
		Save(ctx)

	if err != nil {
		panic(err)
	}
	return p
}
