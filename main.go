package main

// _ is syntax for "import the Postgres driver only for its side effects"
import _ "github.com/lib/pq"

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/PwnySQL/bloggator/internal/config"
	"github.com/PwnySQL/bloggator/internal/database"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Error: Not enough arguments provided. Expect at least two.")
		os.Exit(1)
	}

	cfg, err := config.Read()
	if err != nil {
		fmt.Printf("Error while reading config: %v\n", err)
		os.Exit(1)
	}
	s := state{cfg: &cfg}

	db, err := sql.Open("postgres", cfg.DbUrl)
	s.db = database.New(db)

	cmds := commands{cmds: make(map[string]func(*state, command) error)}
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)

	cmd := command{name: os.Args[1], arguments: os.Args[2:]}
	if err := cmds.run(&s, cmd); err != nil {
		fmt.Printf("Error while executing %s: %v\n", cmd.name, err)
		os.Exit(1)
	}
	os.Exit(0)
}
