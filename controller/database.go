package database

import (
	"log"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	. "github.com/restapi/model"
)

type BooksDB struct {
	Server   string
	Database string
}

var db *mgo.Database

const COLLECTION = "books"


func (m *BooksDB) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

func (m *BooksDB) FindAll() ([]Books, error) {
	var books []Books
	err := db.C(COLLECTION).Find(bson.M{}).All(&books)
	return books, err
}

func (m *BooksDB) FindById(id string) (Books, error) {
	var book Books
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&book)
	return book, err
}

func (m *BooksDB) Insert(book Books) error {
	err := db.C(COLLECTION).Insert(&book)
	return err
}

func (m *BooksDB) Delete(book Books) error {
	err := db.C(COLLECTION).Remove(&book)
	return err
}

func (m *BooksDB) Update(book Books) error {
	err := db.C(COLLECTION).UpdateId(book.ID, &book)
	return err
}