package models

import (
	//"github.com/davecgh/go-spew/spew"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Example is the example object
type Example struct {
	ID    bson.ObjectId `bson:"_id" json:"id"`
	Title string        `bson:"title" json:"title"`
	Body  string        `bson:"body" json:"body"`
}

// FindExample searches for an example based on its id
func FindExample(id string, db *mgo.Database) (*Example, error) {
	ex := Example{}
	bid := bson.ObjectIdHex(id)
	err := db.C("examples").Find(bson.M{"_id": bid}).One(&ex)
	if err != nil {
		return nil, err
	}
	return &ex, nil
}

// ListExamples shows all examples in the database
func ListExamples(db *mgo.Database) (*[]Example, error) {
	var exs []Example
	err := db.C("examples").Find(bson.M{}).All(&exs)
	if err != nil {
		return nil, err
	}
	return &exs, nil
}

// Update is a method on example that updates the copy in the db
func (e Example) Update(db *mgo.Database) error {
	err := db.C("examples").UpdateId(e.ID, e)
	return err
}

// Insert is a method on example that inserts a new example into the db
func (e Example) Insert(db *mgo.Database) error {
	err := db.C("examples").Insert(e)
	return err
}

// Delete is a method on example that will delete the example
func (e Example) Delete(db *mgo.Database) error {
	err := db.C("examples").RemoveId(e.ID)
	return err
}
