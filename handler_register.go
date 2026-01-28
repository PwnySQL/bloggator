package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/PwnySQL/bloggator/internal/database"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.arguments) != 1 {
		return fmt.Errorf("%s expects the user to register as argument", cmd.name)
	}
	name := cmd.arguments[0]
	user, err := s.db.GetUser(context.Background(), name)
	if err == nil {
		return fmt.Errorf("user '%s' already exists. You cannot register the same user.", name)
	}
	fmt.Printf("Register user to %s\n", name)
	user, err = s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID: uuid.New(), CreatedAt: time.Now(), UpdatedAt: time.Now(), Name: name})
	if err != nil {
		return err
	}
	err = s.cfg.SetUser(user.Name)

	fmt.Printf("User %s was registered (at: %v, id: %v)", user.Name, user.CreatedAt, user.ID)

	return err
}
