package main

//Declared types for map of function calls made by user, must return string and take no args
type vanilla func() string

func CmdLookup(cmdLineInput string) string {

	//Defined Map data structure of strings using type vanilla
	//Key, Value similar to dictionary can have more added at a later time
	cmdList := map[string]vanilla{
		"ifconfig": Ifconfig,
		"ps":       Ps,
		"ls":       Ls,
		"pwd":	    Pwd,
	}
	
	//Checks the input to the map to see if a function needs to be called
	for cmds := range cmdList {
		if cmdLineInput == cmds {
			//actual function call happens here using map, not user input
			output := cmdList[cmds]()
			return output
		}
	}
	return "Somethings wrong\n"
}
