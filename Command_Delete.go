package main

import (
	"fmt"
	"bufio"
	"os"
)

func delete(cmd []string){

	args,r := parseArguments(cmd,1)

	if r==2 {
		fmt.Println("[Error] You must specify a key to delete")
		return
	} else if r==3 {
		fmt.Println("[Error] Couldn't parse arguments")
		return
	}

	path := stringToPath(args[0],currentBucket)
	trgt := path[len(path)-1]

	nb := bckt{path[:len(path)-1]}

	if !nb.exists(){
		fmt.Println("[Error] The specified bucket doesn't exist")
		return
	}

	val,suc := nb.getOne(trgt)

	if !suc {
		fmt.Printf("Value for key %s is undefined in %s\n",args[0],nb.bucketString())
		return
	}

	if val.isBucket(){
		fmt.Printf("Are you sure you wish to delete the BUCKET %s?\n",val.key())
	} else {
		fmt.Printf("Are you sure you wish to delete the KEY/VALUE %s?\n",val.key())
	}

	for{
		fmt.Println("Type 'yes' to continue or 'no' to cancel")
		scan := bufio.NewScanner(os.Stdin)
		scan.Scan()
		if scan.Text() == "yes"{
			nb.delete([]byte(trgt))
			fmt.Printf("%s was deleted",val.key())
			break
		} else if scan.Text() == "no" {
			fmt.Println("The database was not changed")
			break
		}
		fmt.Print("Unknown input. ")
	}

}
