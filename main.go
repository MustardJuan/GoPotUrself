package main

import "fmt"
import "os"
import "bufio"

func main() {
	cmdList := make([]string,4)
	cmdList[0] = "ls"
	cmdList[1] = "pwd"
	cmdList[2] = "ifconfig"
	cmdList[3] = "ps"
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		cmdLineInput := scanner.Text()
		for cmds := range cmdList {
			if  cmdLineInput == cmdList[cmds]{
				fmt.Println("call correct cmd function")
			}
		}
	}

	if scanner.Err() != nil {
		fmt.Println("Oops dawg")
	}
}
