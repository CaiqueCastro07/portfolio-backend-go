package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Collection *mongo.Collection
var Ctx = context.TODO()

func InitDatabase() {

	if len(os.Getenv("DB")) == 0 {
		log.Fatal("Configure a variável de ambiente com a uri do mongoDB antes de iniciar.")
	}

	clientOptions := options.Client().ApplyURI(os.Getenv("DB"))
	client, err := mongo.Connect(Ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(Ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	Collection = client.Database("portfolio").Collection("messages")

	app := &cli.App{
		Name:     "tasker",
		Usage:    "A simple CLI program to manage your tasks",
		Commands: []*cli.Command{},
	}

	err2 := app.Run(os.Args)
	if err2 != nil {
		log.Fatal(err2)
	}

	fmt.Println("Banco de dados conectado com sucesso.")
}
