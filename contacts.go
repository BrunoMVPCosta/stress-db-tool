package main

import (
	"fmt"
	"strconv"

	"github.com/dmgk/faker"
	"gopkg.in/mgo.v2/bson"
)

// Contact represents a contact in mongodb
type Contact struct {
	ID      bson.ObjectId `bson:"_id,omitempty"`
	Name    string        `bson:"name"`
	Points  int32         `bson:"points"`
	Padding string        `bson:"padding"`
}

// GenerateMany is used to generate a given number of contacts
func GenerateMany(numberOfContacts int) (contacts []Contact) {
	fmt.Printf("Generating %d contacts\n", numberOfContacts)
	for index := 0; index < numberOfContacts; index++ {
		value := faker.Number().Between(1, 1000000)

		points, err := strconv.ParseInt(value, 10, 32)
		if err != nil {
			fmt.Println("Error converting to integer: ", err)
		}

		contacts = append(contacts, Contact{
			Name:    faker.Name().Name(),
			Points:  int32(points),
			Padding: faker.Code().Isbn13(),
		})
	}

	return contacts
}

func generate() (contact Contact) {
	value := faker.Number().Between(1, 1000000)

	points, err := strconv.ParseInt(value, 10, 32)
	if err != nil {
		fmt.Println("Error converting to integer: ", err)
	}

	contact = Contact{
		Name:    faker.Name().Name(),
		Points:  int32(points),
		Padding: faker.Code().Isbn13(),
	}

	return contact
}
