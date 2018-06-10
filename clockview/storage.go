package clockskew

import (
	"gopkg.in/mgo.v2"
	"log"
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"math"
)

func MongoStorage(){
	session, err := mgo.Dial("localhost")
	
	if err != nil {
       panic(err)
	}
	
	defer session.Close()
	
	defer close(ClockChannel)


	for{
		cs := <- ClockChannel
		c := session.DB( "d").C("clocks")

		// placeholder
		var clocks []Clock
        // get all clocks with this host ip and sort by the localts
		err = c.Find(bson.M{"hostip": cs.HostIp}).Sort("localts").All(&clocks)
        /// see the results
		fmt.Println("Results All: ", clocks)
		// initial value of previous ts is 0
		var previousTs = uint32(0)
		// the amount offset we allow
		var allowedOffset uint32 = uint32(50000);
        // take every clock in the list of clocks
		for _, clock := range clocks {
			// set the current clock timestamp
			currentTs := clock.HostTs
			// if the previous ts is 0 then set the first timestamp
			if (previousTs == 0) {
				previousTs = currentTs
			} else {
				// if we have a previous value to measure we check it's difference
				// taking an absolute value
				offset := math.Abs(float64(previousTs) - float64(currentTs))
				// if tht difference is greater than we allow, it should cry and flag tge record
				if (uint32(offset) > allowedOffset) {
					fmt.Printf(
						"an offset was" +
							   "observed greater " +
							    "than allowed" +
								"it was off by " +
								string(uint32(offset)) +
								" units")
					err = c.Insert(cs)
				} else {
					err = c.Insert(cs)
				}
			}

		}
		if err != nil {
			log.Fatal(err)
		}
	}
}


