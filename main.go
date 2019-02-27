package main

import "fmt"
import "os"
import "bufio"

//Declared types for map of function calls made by user, must return string and take no args
type vanilla func() string

func main() {
	//Defined Map data structure of strings using type vanilla
	//Key, Value similar to dictionary can have more added at a later time
	cmdList := map[string] vanilla {
	"ifconfig": ifconfig,
	}
	//Scanner class to indefinetly read in input from user
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("guestUser@HomeOffice:~$")
	for scanner.Scan() {
		fmt.Print("guestUser@HomeOffice:~$")
		cmdLineInput := scanner.Text()
		//Checks the input to the map to see if a function needs to be called
		for cmds := range cmdList {
			if cmdLineInput == cmds{
				fmt.Println("That's a command my dude")
				fmt.Println(cmdList[cmds])
			}
		}
	}
	//error checking for the scanner
	if scanner.Err() != nil {
		fmt.Println("Oops dawg")
	}
}
