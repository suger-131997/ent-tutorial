package main

import (
	"context"
	"ent-tutorial/ent"
	"ent-tutorial/ent/pet"
	"ent-tutorial/ent/user"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	client, err := ent.Open("mysql", "docker:docker@tcp(localhost:3306)/test_database?parseTime=True")
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}
	defer client.Close()

	ctx := context.Background()
	user1, _ := client.User.Query().Where(user.NameEQ("user1")).Only(ctx)
	fmt.Println(user1)

	pets, _ := user1.QueryPets().All(ctx)
	fmt.Println(pets)

	userX, _ := client.User.Query().Where(user.HasPetsWith(pet.AgeEQ(1))).Only(ctx)
	fmt.Println(userX)
}
