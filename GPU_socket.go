package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"math/rand"
	"time"
	"strconv"
	"github.com/GoPotUrself/shell"
)

func main() {
	
	ipAddress := os.Args[1]
	port := ":8080"
	if len(os.Args) > 2 {
		port = ":" + os.Args[2]
	}

	//connect to this socket
	address := ipAddress + port
	conn, _ := net.Dial("tcp", address)
	
	currentTime := time.Now()
	month := currentTime.Month()
	weekday := currentTime.Weekday()
	year := strconv.Itoa(currentTime.Year())
	dayOut := weekday.String()[0:3]
	monthOut := month.String()[0:3]
	timeOut := currentTime.String()[11:19]

	output := "Linux f6651192effa 4.9.125-linuxkit #1 SMP " + 
		dayOut + 
		" " + 
		monthOut + 
		" " + 
		weekday.String() + 
		" " + 
		timeOut + 
		" UTC " + 
		year + 
		" x86_64 x86_64 x86_64 GNU/Linux\n" +
		time.Unix(rand.Int63n(24),rand.Int63n(60)).String() + 
		" up 15 min, 1 users, load average: 0.00, 0.02, 0.00\n" + 
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
		
		//output = shell.CmdLookup(output) + "$ "
		fmt.Fprintf(conn, output)
	}
}
