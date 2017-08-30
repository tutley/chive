package main

import (
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type example struct {
	id    bson.ObjectID `bson:"_id" json:"id"`
	title string        `bson:"title" json:"title"`
	body  string        `bson:"body" json:"body"`
}

// find searches for an example based on its id
func findExample(string title, db *mgo.Database) (*example, error) {
	ex := example{}
	err := db.C("examples").Find(bson.M{"title": title}).One(&ex)
	if err != nil {
		return nil, err
	}
	return &ex, nil
}

// list shows all examples in the database
func listExamples(db *mgo.Database) (*[]example, error) {
	var exs []example
	err := db.C("examples").Find(bson.M{}).All(&exs)
	if err != nil {
		return nil, err
	}
	return &exs, nil
}

// update is a method on example that updates the copy in the db
func (e example) update(db *mgo.Database) error {
	q := bson.M{"_id": e.id}
	up := bson.M{"$set": bson.M{"title": e.title, "body": e.body}}
	err := db.C("examples").Update(q, up)
	return err
}
