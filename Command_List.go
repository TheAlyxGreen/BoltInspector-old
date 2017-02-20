package main

import (
	"fmt"
)

func list(cmd []string) {
	if !currentBucket.exists(){
		fmt.Println("[Error] Bucket does not exist. Returning to root...")
		currentBucket.reset()
		return
	}
	show := 0 // 0=all,1=buckets,2=keys
	verbose := false

	args,r := parseArguments(cmd,0)

	if r==3 {
		fmt.Println("[Error] Couldn't parse arguments")
		return
	} else if r==0 {
		for i:=0;i<len(args);i++{
			if args[i]=="b" || args[i]=="-b"{
				show=1
			} else if args[i]=="k" || args[i]=="-k" {
				show = 2
			} else if args[i]=="v" || args[i]=="-v" {
				verbose=true
			}
		}
	}

	bckts,keys := currentBucket.getAllSeparated()
	rpath := "./"
	if currentBucket.isRoot(){
		rpath = "~/"
	}
	if verbose {
		rpath = ""
	}
	if show==0||show==1 {
		for _,val := range bckts{
			fmt.Println(rpath+val.toString(verbose))
		}
	}
	if show==0||show==2 {
		for _,val := range keys{
			fmt.Println(val.toString(verbose))
		}
	}
}