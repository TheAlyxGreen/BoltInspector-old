package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
)

func delete(cmd []string){

	if len(cmd)==1{
		fmt.Println("[Error] You must specify a key to delete.")
		return
	}

	args := strings.Split(cmd[1]," ")

	val,suc := currentBucket.getOne(args[0])

	if !suc {
		fmt.Println("Value for key \"" + args[0] + "\" is undefined in this bucket (" + currentBucket.bucketString() + ")")
		return
	}

	if val.isBucket(){
		fmt.Print("Are you sure you wish to delete the BUCKET ")
	} else {
		fmt.Print("Are you sure you wish to delete the KEY/VALUE ")
	}
	fmt.Println(val.key()+"?")

	for{
		fmt.Println("Type 'yes' to continue or 'no' to cancel")
		scan := bufio.NewScanner(os.Stdin)
		scan.Scan()
		if scan.Text() == "yes"{
			currentBucket.delete([]byte(args[0]))
			fmt.Println("The specified key was deleted")
			break
		} else if scan.Text() == "no" {
			fmt.Println("The database was not changed")
			break
		}
		fmt.Print("Unknown input. ")
	}

}
