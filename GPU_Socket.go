package main

import "net"
import "fmt"
import "bufio"
import "time"
import "rand"
import "strings"

/*
So a few floating things here:
1. Variables to be made and taken as arguments
	-NEED to take IP from session, dont forget about XFF
	-Want to take PORT from php file to open port as specified
2. From the PHP we need a way to make this outbound connection loop
	-I think we can handle this with concurrency in golang
	-Or maybe theres some php that we can use to run the code on page visit
*/

func main() {
	
	
	//connect to this socket
	conn, _ := net.Dial("tcp", "127.0.0.1:8080")
	
	/*
	Displays the initial reverse shell, improvements:
	1.Formatting isn't perfect, time is awkward
	2. Fixed the random time generator, maybe few hours behind normal time? 
	3. Keep displaying dollar sign, should just append to output at bottom of loop
	*/
	currentTime := time.Now()
	month := currentTime.Month()
	weekday := currentTime.Weekday()
	dayOut := weekday.String()[0:3]
	monthOut := month.String()[0:3]

	output := "Linux f6651192effa 4.9.125-linuxkit #1 SMP " + dayOut + " " + monthOut + " " + 
	weekday.String() + " 08:20:28 UTC 2018 x86_64 x86_64 x86_64 GNU/Linux\n" +
	time.Unix(rand.Int63n(24),rand.Int63n(60)).String() + " up 10 min, 0 users, load average: 0.00, 0.02, 0.00\n" + 
	"USER	TTY	FROM		LOGIN@	IDLE	JCPU	PCPU	WHAT\n" + 
	"uid=33(www-data) gid=33(www-data) groups=33(www-data)\n" +
	"/bin/sh: 0: can't access tty; job control turned off\n" + 
	"$ " 
    
	// send to socket
	fmt.Fprintf(conn, output)
	
	for { 
    
		// listen for reply and then display it to the web app owner
		message, err := bufio.NewReader(conn).ReadString('\n')
		
		if err != nil {

			fmt.Println(err)
			return
		}
		
		output := strings.TrimSpace(string(message))
		fmt.Print("Message from server: " + message)
		
		// send the fake output to the malicious server
		output = CmdLookup(output)
		fmt.Fprintf(conn, output)

	}
}
