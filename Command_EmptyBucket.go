package main

import (
	"fmt"
	"bufio"
	"os"
)

func emptyBucket(cmd []string) {

	args,r := parseArguments(cmd,0)

	bp := currentBucket.path

	if r==0 {
		bp = stringToPath(args[0],currentBucket)
	} else if r==3 {
		fmt.Println("[Error] Couldn't parse arguments")
		return
	}

	nb := bckt{bp}

	if !nb.exists(){
		fmt.Println("The specified bucket does not exist")
		return
	}

	bc,vc := nb.countBoth()

	fmt.Printf("Are you sure you want to delete all %d %s and the %d %s in %s?\n",vc,valPlural(vc),bc,bcktPlural(bc),nb.bucketString())

	for{
		fmt.Println("Type 'yes' to continue or 'no' to cancel")
		scan := bufio.NewScanner(os.Stdin)
		scan.Scan()
		if scan.Text() == "yes"{
			nb.empty()
			fmt.Println("The specified bucket was emptied")
			break
		} else if scan.Text() == "no" {
			fmt.Println("The database was not changed")
			break
		}
		fmt.Print("Unknown input. ")
	}
}