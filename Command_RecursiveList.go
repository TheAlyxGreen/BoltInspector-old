package main

import (
	"strings"
	"fmt"
	"regexp"
	"strconv"
)

func rlist(cmd []string) {
	if !currentBucket.exists(){
		fmt.Println("[Error] Bucket does not exist. Returning to root...")
		currentBucket.reset()
		return
	}
	verbose := false
	maxRecurse := 3

	args,r := parseArguments(cmd,0)

	if r==3 {
		fmt.Println("[Error] Couldn't parse arguments")
		return
	} else if r==0 {
		for i:=0;i<len(args);i++{
			if args[i]=="v" || args[i]=="-v" {
				verbose=true
			} else if strings.HasPrefix(args[i],"d=")||strings.HasPrefix(args[i],"depth="){
				re := regexp.MustCompile("[^\\d]")
				tmp := re.ReplaceAllString(args[i],"")
				in,er:=strconv.Atoi(tmp)
				if er!= nil{
					fmt.Println("[Error] Invalid length passed")
					return
				}
				maxRecurse=in
			}
		}
	}

	recurseDump(currentBucket,maxRecurse,0,verbose)

}

func recurseDump(b bckt, rMax int, rCur int, verbose bool){
	bckts,vals := b.getAllSeparated()
	rpath := "."
	if currentBucket.isRoot(){
		rpath = "~"
	}
	for i,val := range b.path{
		if i>=len(currentBucket.path){
			rpath=rpath+"/"+val
		}
	}
	for _,val := range vals{
		if !verbose{fmt.Print(rpath+"/")}
		fmt.Println(val.toString(verbose))
	}
	for _,val := range bckts{
		if !verbose{fmt.Print(rpath+"/")}
		fmt.Print(val.toString(verbose))
		if verbose && rMax<=rCur {
			fmt.Printf("- Contents Outside Recurse Depth (Depth = %d)\n\n", rMax)
		} else if rMax<=rCur {
			fmt.Println()
		} else if rMax>rCur {
			fmt.Println()
			recurseDump(val.asBucket(), rMax, rCur+1, verbose)
		}
	}
}