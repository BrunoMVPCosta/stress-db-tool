package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Connection struct {
	Session *mgo.Session
}

func main() {

	configuration := ReadConfiguration("conf.json")
	configuration.PrintConfiguration()

	session, err := mgo.Dial(configuration.MongoConnectionString)
	if err != nil {
		log.Fatal(err)
	}
	session.SetMode(mgo.SecondaryPreferred, false)
	defer session.Clone()

	var numberOfContacts = 30000
	var waitGroup sync.WaitGroup

	quit := make(chan bool)
	channelPrintStats := make(chan bool)
	go func() {
		var queryNumber = 1
		r := rand.New(rand.NewSource(99))

		channelStats := make(chan int, 10)
		go RegisterStats(channelStats, channelPrintStats)

		concurrency := 15
		semaphone := make(chan bool, concurrency)

		for {
			semaphone <- true
			select {
			case <-quit:
				return
			default:
				points := r.Int31n(1000000)
				go func() {
					defer func() { <-semaphone }()
					RunQuery(channelStats, queryNumber, session, points)
				}()
			}
			queryNumber++
			time.Sleep(2 * time.Millisecond)
		}
	}()

	var contacts = GenerateMany(numberOfContacts)

	if configuration.InsertMany {
		waitGroup.Add(1)
		go func() {
			InsertMany(contacts, session)
			waitGroup.Done()
		}()
		waitGroup.Wait()
		time.Sleep(5 * time.Second)
	}

	if configuration.BulkInsertWithSleep {
		waitGroup.Add(1)
		go func() {
			BulkInsertWithSleeps(contacts, session)
			waitGroup.Done()
		}()
		waitGroup.Wait()
		channelPrintStats <- true
		time.Sleep(5 * time.Second)
	}

	if configuration.BulkInsert {
		waitGroup.Add(1)
		go func() {
			BulkInsert(contacts, session)
			waitGroup.Done()
		}()
		waitGroup.Wait()
		channelPrintStats <- true
		time.Sleep(5 * time.Second)
	}

	waitGroup.Add(1)
	go func() {
		time.Sleep(60 * time.Second)
		waitGroup.Done()
	}()
	waitGroup.Wait()
	channelPrintStats <- true
	time.Sleep(2000)
	quit <- true
	fmt.Printf("All done... now leaving!")
}

// RunQuery is where the query to mongo is done
func RunQuery(channelStats chan int, queryNumber int, session *mgo.Session, points int32) {

	start := time.Now()

	sessionCopy := session.Copy()
	defer sessionCopy.Close()

	customers := sessionCopy.DB("customer").C("customer")

	_, err := customers.Find(bson.M{
		"points": bson.M{"$lte": points},
	}).Count()

	if err != nil {
		log.Fatal(err)
	}

	elapsed := time.Since(start)

	channelStats <- int(elapsed)
}
