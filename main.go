package main

import (
	"context"

	_ "github.com/joho/godotenv/autoload"
	"github.com/wildanfaz/go-oauth2/cmd"
)

func main() {
	cmd.InitCmd(context.Background())
}
