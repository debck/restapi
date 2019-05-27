package database

import (
	"log"
	"github.com/restapi/model"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Books struct {
	Server   string
	Database string
}

var db *mgo.Database

const COLLECTION = "books"


func (m *Books) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

func (m *Books) FindAll() ([]Books, error) {
	var books []Books
	err := db.C(COLLECTION).Find(bson.M{}).All(&books)
	return books, err
}

func (m *Books) FindById(id string) (Books, error) {
	var book Books
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&book)
	return book, err
}

func (m *Books) Insert(book Books) error {
	err := db.C(COLLECTION).Insert(&book)
	return err
}

func (m *Books) Delete(book Books) error {
	err := db.C(COLLECTION).Remove(&book)
	return err
}

func (m *Books) Update(book Books) error {
	err := db.C(COLLECTION).UpdateId(book.ID, &book)
	return err
}