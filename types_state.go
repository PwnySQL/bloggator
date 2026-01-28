package main

import (
	"github.com/PwnySQL/bloggator/internal/config"
	"github.com/PwnySQL/bloggator/internal/database"
)

type state struct {
	cfg *config.Config
	db  *database.Queries
}
