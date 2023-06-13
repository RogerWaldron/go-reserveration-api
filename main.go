package main

import (
	"context"
	"flag"
	"log"

	"github.com/RogerWaldron/go-reserveration-api/api"
	"github.com/RogerWaldron/go-reserveration-api/db"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dbURI = "mongodb://localhost:27017"

var config = fiber.Config{
	ErrorHandler: func(c *fiber.Ctx, err error) error {
		return c.JSON(map[string]string{"error": err.Error()})
	},
}

func main() {
	listenAddress := flag.String("listenAddress", ":3200", "This listen address for API server")
	flag.Parse()
 
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dbURI))
	if err != nil {
		log.Fatal(err)
	}
	userHandler := api.NewUserHandler(db.NewMongoUserStore(client))
	
	app := fiber.New(config)
	appv1 := app.Group("/api/v1")
	
	appv1.Post("/user", userHandler.HandlePostUser)

	appv1.Get("/user", userHandler.HandleGetUsers)
	appv1.Get("/user/:id", userHandler.HandleGetUserByID)

	app.Listen(*listenAddress) 
}