package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http/httptest"
	"testing"

	"github.com/RogerWaldron/go-reserveration-api/db"
	"github.com/RogerWaldron/go-reserveration-api/types"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const testURI = "mongodb://localhost:27017"

type testDB struct {
	db.UserStore
}

func (tdb *testDB) teardown(t *testing.T) {
	err := tdb.UserStore.Drop(context.TODO())
	if err != nil {
		t.Fatal(err)
	}
}

func setup(t *testing.T) *testDB {

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(testURI))
	if err != nil {
		log.Fatal(err)
	}

	return &testDB{
		UserStore: db.NewMongoUserStore(client, db.DB_TEST_NAME),
	}
}

 
func TestPostUser(t *testing.T) {
	tDB := setup(t)
	defer tDB.teardown(t)

	app := fiber.New()
	
	userHandler := NewUserHandler(tDB.UserStore)
	app.Post("/", userHandler.HandlePostUser)

	params := types.CreateUserParams{
		Email:     "some@foo.com",
		FirstName: "James",
		LastName:  "Foo",
		Password:  "lkdfjkdsjfklfdjkedf",
	}

	b, _ := json.Marshal(params)
	req := httptest.NewRequest("POST", "/", bytes.NewReader(b))
	req.Header.Add("Content-Type", "application/json")

	// without 3 sec timeout test fails
	resp, err := app.Test(req, 3000) 
	if err != nil {
		t.Error(err)
	}

	var user types.User
	err = json.NewDecoder(resp.Body).Decode(&user)
	if err != nil {
		t.Error(err)
	}
	if len(user.ID) == 0 {
		t.Error("expected a user ID")
	}
	fmt.Printf("%+v", user)
	if user.FirstName != params.FirstName {
		t.Errorf("expected firstName %s but got %s", params.FirstName, user.FirstName)
	}
	if user.LastName != params.LastName {
		t.Errorf("expected lastName %s but got %s", params.LastName, user.LastName)
	}
	if user.Email != params.Email {
		t.Errorf("expected email %s but got %s", params.Email, user.Email)
	}
	if len(user.EncryptedPassword) > 0 {
		t.Error("expected  EncryptedPassword not to be returned")
	} 
} 