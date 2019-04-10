package main

import "net"
import "fmt"
import "bufio"

/*
So a few floating things here:
1. Variables to be made and taken as arguments
	-NEED to take IP from session, dont forget about XFF
	-Want to take PORT from php file to open port as specified
2. From the PHP we need a way to make this outbound connection loop
	-I think we can handle this with concurrency in golang
	-Or maybe theres some php that we can use to run the code on page visit
3. For some reason CmdLookup isnt working now but I'm not entirely sure why?
	-Could be something I'm doing wrong but the message sent is a string
	-Need to figure out whats wrong there
*/

func main() {
	
	
	//connect to this socket
	conn, _ := net.Dial("tcp", "127.0.0.1:8080")
	
	//read in input from stdin
	output := "Hi, yes, hello, You have a reverse shell MY dude"
    
	// send to socket
	fmt.Fprintf(conn, output + "\n")
	
	for { 
    
		// listen for reply and then display it to the web app owner
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Message from server: " + message)
		
		// send the fake output to the malicious server
		output := CmdLookup(message)
		fmt.Fprintf(conn, output)

	}
}
