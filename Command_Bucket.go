package main

import (
	"fmt"
)

func bucket(cmd []string){

	args,r := parseArguments(cmd,1)

	if r==2{
		fmt.Println("[Error] You must specify a key for the bucket.")
		return
	} else if r==3{
		fmt.Println("[Error] Couldn't parse arguments.")
		return
	}

	_,suc := currentBucket.getOne(args[0])

	if suc {
		fmt.Println("Value for key \"" + args[0] + "\" is already defined in this bucket (" + currentBucket.bucketString() + ")")
		return
	}

	currentBucket.insertBucket([]byte(args[0]))

}