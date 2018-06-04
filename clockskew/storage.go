package clockskew

import (
	"gopkg.in/mgo.v2"
	"log"
)

func MongoStorage(){
	session, err := mgo.Dial("localhost")
	
	if err != nil {
       panic(err)
	}
	
	defer session.Close()
	
	defer close(ClockSkewChannel)
	
    c := session.DB( "clock").C("skews")

	for{
		cs := <- ClockSkewChannel
		

		err = c.Insert(cs)
		
		if err != nil {
			log.Fatal(err)
		}
	}
}
