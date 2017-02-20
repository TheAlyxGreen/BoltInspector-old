package main

import (
	"fmt"
)

func cd(cmd []string){

	args,r := parseArguments(cmd,1)

	if r==2{
		fmt.Println("[Error] You must enter a destination")
	}

	bp := stringToPath(args[0],root)

	nb := bckt{bp}

	if nb.exists(){
		currentBucket.path=bp
	} else {
		fmt.Println("[Error] Invalid bucket path: " + nb.bucketString())
	}

}