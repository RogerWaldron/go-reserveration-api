package db

import "context"
 
const DB_NAME = "reservation"
const DB_TEST_NAME = DB_NAME + "_test"

type Dropper interface {
	Drop(context.Context) error
}

// "any" so can pass in mongoDB ID which is a primitive Object
type Map map[string]any