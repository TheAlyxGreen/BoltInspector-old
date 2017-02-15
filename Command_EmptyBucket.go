package main

import (
	"strings"
	"fmt"
	"bufio"
	"os"
	"strconv"
)

func emptyBucket(cmd []string) {

	target:="~"

	if len(cmd)>1{
		target = cmd[1]
	}

	bp := []string{}

	if target=="~"{
		bp = []string{"~"}
	} else if strings.HasPrefix(target,"./"){

		// if the path is relative
		bp = currentBucket.path
		target = strings.Replace(target,"./","",1)
		tmp := escapedSplit(target)
		for i:=0;i<len(tmp);i++{
			bp = append(bp,tmp[i] )
		}
	} else if strings.HasPrefix(target,"/")||strings.HasPrefix(target,"~/"){

		// if the path is absolute
		if strings.HasPrefix(target,"/"){
			target = "~"+target
		}
		tmp := escapedSplit(target)
		for i:=0;i<len(tmp);i++{
			bp = append(bp,tmp[i] )
		}
	} else {

		// assume the path is relative
		bp = currentBucket.path
		tmp := escapedSplit(target)
		for i:=0;i<len(tmp);i++{
			bp = append(bp,tmp[i] )
		}
	}

	nb := bckt{bp}

	if !nb.exists(){
		fmt.Println("The specified bucket does not exist")
		return
	}

	vals := nb.getAll()
	keyCount := 0
	BktCount := 0
	for i:=0;i<len(vals);i++{
		if vals[i].isBucket(){
			BktCount++
		} else {
			keyCount++
		}
	}

	fmt.Println("Are you sure you want to delete all "+strconv.Itoa(keyCount)+" values and all "+strconv.Itoa(BktCount)+" buckets in ("+nb.bucketString()+")?")


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