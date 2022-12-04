package main

import (
	"context"
	"ent-tutorial/ent"
	"ent-tutorial/ent/pet"
	"ent-tutorial/ent/user"
	"entgo.io/ent/dialect/sql"
	"fmt"
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
	user1, _ := client.User.Query().Where(user.Name("user1")).Only(ctx)
	fmt.Println(user1)

	user2, _ := client.User.Query().Where(
		user.And(
			user.NameContains("user"),
			user.NameContains("2"),
		)).Only(ctx)
	fmt.Println(user2)

	pets, _ := user1.QueryPets().All(ctx)
	fmt.Println(pets)

	userX, _ := client.User.Query().Where(user.HasPetsWith(pet.Age(1))).Only(ctx)
	fmt.Println(userX)

	userY, _ := client.User.Create().SetName("userY").SetRegisteredAt(time.Now()).AddPets(pets...).Save(ctx)
	fmt.Println(userY)
	petX, _ := client.Pet.Create().SetName("pet").SetAge(2).SetOwner(userX).Save(ctx)
	fmt.Println(petX)

	petX, _ = petX.Update().SetName("petX").Save(ctx)
	fmt.Println(petX)

	users, _ := client.User.Query().
		Order(ent.Asc(user.FieldName)).
		Offset(2).
		Limit(2).
		All(ctx)
	fmt.Println(users)

	pets, _ = client.Pet.Query().
		Order(func(s *sql.Selector) {
			t := sql.Table(user.Table)
			s.Join(t).On(s.C(pet.OwnerColumn), t.C(user.FieldID))
			s.OrderBy(t.C(user.FieldName))
		}).
		Unique(false).
		All(ctx)
	fmt.Println(pets)
}
