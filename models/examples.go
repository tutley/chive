package models

import (
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
func FindExample(title string, db *mgo.Database) (*Example, error) {
	ex := Example{}
	err := db.C("examples").Find(bson.M{"title": title}).One(&ex)
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
	q := bson.M{"_id": e.ID}
	up := bson.M{"$set": bson.M{"title": e.Title, "body": e.Body}}
	err := db.C("examples").Update(q, up)
	return err
}
