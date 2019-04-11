package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/GoPotUrself/shell"
)

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
	address := os.Args[1] + ":8080"
	conn, _ := net.Dial("tcp", address)

	//read in input from stdin
	output := "Hi, yes, hello, You have a reverse shell MY dude"

	// send to socket
	fmt.Fprintf(conn, output+"\n")

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
		output = shell.CmdLookup(output)
		fmt.Fprintf(conn, output)

	}
}
