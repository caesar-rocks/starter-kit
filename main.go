package main

import (
	"fmt"
	"os"
	"starter_kit/config"
	"starter_kit/database/migrations"
	"strings"

	"github.com/caesar-rocks/orm"
)

const (
	possibleArgs = "--serve --migrate --rollback --reset --seed"
)

func main() {
	args := os.Args
	if len(args) != 2 || !strings.Contains(possibleArgs, args[1]) {
		fmt.Println("Usage: go run caesar.go [--serve | --migrate | --rollback | --reset | --seed]")
		os.Exit(1)
	}

	switch args[1] {
	case "--serve":
		env := config.ProvideEnvironmentVariables()
		config.ProvideApp(env).Run()
	case "--migrate":
		getDB().Migrate(migrations.Migrations)
	case "--rollback":
		getDB().Rollback(migrations.Migrations)
	case "--reset":
		getDB().Reset(migrations.Migrations)
	case "--seed":
		getDB().Seed()
	}
}

func getDB() *orm.Database {
	env := config.ProvideEnvironmentVariables()
	db := config.ProvideDatabase(env)
	return db
}
