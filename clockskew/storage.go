package clockskew

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"log"
)



func Storage(){
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	defer StorageFile.Close()
	defer close(ClockSkewChannel)

	for{
		cs := <- ClockSkewChannel
		c := session.DB("clock").C("skews")
		err = c.Insert(cs)
		if err != nil {
			log.Fatal(err)
		}
		item := fmt.Sprintf("%d %s %d %d\n",cs.Clock, cs.Taddr, cs.SrcTS, cs.Skew)
		StorageFile.WriteString(item) 
	}
}
