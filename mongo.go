package main

import (
	"fmt"
	"time"

	mgo "gopkg.in/mgo.v2"
)

// BulkInsert inserts contacts into mongodb using bulk apis
func BulkInsert(contacts []Contact, session *mgo.Session) {
	fmt.Println("[BulkInsert] Bulk insert contacts")

	sessionCopy := session.Copy()
	defer sessionCopy.Close()

	start := time.Now()

	customers := sessionCopy.DB("customer").C("customer")
	bulk := customers.Bulk()
	for index := 0; index < len(contacts); index++ {
		bulk.Insert(contacts[index])
	}

	_, err := bulk.Run()
	if err != nil {
		fmt.Println(err)
	}

	elapsed := time.Since(start)
	fmt.Printf("[BulkInsert] Inserted %d contacts and it took %s \n", len(contacts), elapsed)
}

// InsertMany will insert many contacts, but without using bulk apis
func InsertMany(contacts []Contact, session *mgo.Session) {
	fmt.Println("[InsertMany] Insert many contacts - no bulk")

	sessionCopy := session.Copy()
	defer sessionCopy.Close()

	start := time.Now()
	customers := sessionCopy.DB("customer").C("customer")

	for index := 0; index < len(contacts); index++ {
		var contact = contacts[index]
		err := customers.Insert(contact)

		if err != nil {
			fmt.Println(err)
		}
	}

	elapsed := time.Since(start)
	fmt.Printf("[InsertMany] Inserted %d contacts without bulk and it took %s \n", len(contacts), elapsed)
}

// BulkInsertWithSleeps inserts contacts into mongodb using bulk apis
func BulkInsertWithSleeps(contacts []Contact, session *mgo.Session) {
	fmt.Println("[BulkInsertWithSleeps] Bulk insert contacts")

	sessionCopy := session.Copy()
	defer sessionCopy.Close()

	start := time.Now()

	customers := sessionCopy.DB("customer").C("customer")

	bulk := customers.Bulk()
	threadhold := 1000
	for index := 0; index < len(contacts); index++ {
		bulk.Insert(contacts[index])
		threadhold--
		if threadhold == 0 {
			_, err := bulk.Run()
			if err != nil {
				fmt.Println(err)
			}
			threadhold = 1000
			time.Sleep(100 * time.Millisecond)
		}
	}

	_, err := bulk.Run()
	if err != nil {
		fmt.Println(err)
	}

	elapsed := time.Since(start)
	fmt.Printf("[BulkInsertWithSleeps] Inserted %d contacts and it took %s \n", len(contacts), elapsed)
}
