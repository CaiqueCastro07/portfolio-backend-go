package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MessagesCollection *mongo.Collection
var TasksCollection *mongo.Collection
var Ctx = context.TODO()

func InitDatabase() {

	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Configure a variável de ambiente .env antes de iniciar.")
	}

	if len(os.Getenv("DB")) == 0 {
		log.Fatal("Configure a variável de ambiente com a uri do mongoDB 'key = DB' antes de iniciar.")
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

	MessagesCollection = client.Database("portfolio").Collection("messages")
	TasksCollection = client.Database("portfolio").Collection("tasks")

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
