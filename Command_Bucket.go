package main

import (
	"fmt"
	"strings"
)

func bucket(cmd []string){
	if len(cmd)==1{
		fmt.Println("[Error] You must specify a key to delete.")
		return
	}

	args := strings.Split(cmd[1]," ")

	_,suc := currentBucket.getOne(args[0])

	if suc {
		fmt.Println("Value for key \"" + args[0] + "\" is already defined in this bucket (" + currentBucket.bucketString() + ")")
		return
	}

	currentBucket.insertBucket([]byte(args[0]))

}