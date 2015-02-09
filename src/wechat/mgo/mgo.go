package mgo

import (
	"fmt"
	"labix.org/v2/mgo"
	// "labix.org/v2/mgo/bson"
)

// type Person struct {
// 	Name  string
// 	Phone string
// }

// func Mgo() {
// 	session, err := mgo.Dial("127.0.0.1")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer session.Close()

// 	// Optional. Switch the session to a monotonic behavior.
// 	session.SetMode(mgo.Monotonic, true)

// 	c := session.DB("test").C("people")
// 	c.Insert(...)
// 	err = c.Insert(&Person{"Ale", "+55 53 8116 9639"},
// 		&Person{"Cla", "+55 53 8402 8510"})
// 	if err != nil {
// 		panic(err)
// 	}

// 	result := Person{}
// 	err = c.Find(bson.M{"name": "Ale"}).One(&result)
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Println("Phone:", result.Phone)
// }

func MongoInsert(v interface{}, collection string) error {
	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		fmt.Println(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("wechat").C(collection)
	err = c.Insert(v)
	if err != nil {
		fmt.Println(err)

	}
	return err
}
