package clockskew

import (
	"fmt"
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
	
    c, err := session.DB(name: "clock").C(name:"skews")
    
    if err != nil { 
    	panic(err)
    }
    
	for{
		cs := <- ClockSkewChannel
		
		c := session.DB("clock").C("skews")
		
		err = c.Insert(cs)
		
		if err != nil {
			log.Fatal(err)
		}
	}
}
