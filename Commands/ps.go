package main

import "fmt"
import "math/rand"
import "time"
import "strconv"

//for now some random value will do for the PID
//However in the future this will be updated to reflect any updates
//e.g. more processes open so the PID should go up for x,y,z services
//additionally the time will be treated the same

func ps() string {

	currentTime := time.Now()
	formatted := fmt.Sprintf("%02d:%02d:%02d", 
	currentTime.Hour(), currentTime.Minute(), currentTime.Second())
	output := ("  PID TTY          TIME CMD\n" +
 	strconv.Itoa(rand.Intn(5000))+" pts/0    "+formatted+" bash\n" +
 	strconv.Itoa(rand.Intn(5000))+" pts/0    "+formatted+" ps\n")

	return output
}
