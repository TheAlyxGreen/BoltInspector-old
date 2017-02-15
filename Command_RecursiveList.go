package main

import (
	"strings"
	"fmt"
	"regexp"
	"strconv"
)

func rlist(cmd []string) {
	verbose := false
	maxRecurse := 3
	var args []string
	if len(cmd)>1{
		args = strings.Split(cmd[1]," ")
		for i:=0;i<len(args);i++{
			if args[i]=="v" || args[i]=="-v" {
				verbose=true
			} else if strings.HasPrefix(args[i],"d=")||strings.HasPrefix(args[i],"depth="){
				re := regexp.MustCompile("[^\\d]")
				tmp := re.ReplaceAllString(args[i],"")
				in,er:=strconv.Atoi(tmp)
				if er!= nil{
					fmt.Println("Invalid length passed")
					return
				}
				maxRecurse=in
			}
		}
	}

	if !currentBucket.exists(){
		fmt.Println("[Error] Bucket does not exist. Returning to root...")
		currentBucket.reset()
		return
	}

	recurseDump(currentBucket,maxRecurse,0,verbose)

}

func recurseDump(b bckt,rMax int, rCur int, verbose bool){
	vals := b.getAll()
	if vals == nil{
		fmt.Println("[Error] Invalid bucket called in Recursive Listing ("+b.bucketString()+").")
		return
	}
	rpath := "."
	for i,val := range b.path{
		if i>len(currentBucket.path){
			rpath=rpath+"/"+val
		}
	}
	for _,val := range vals{
		if val.isBucket() {
			if verbose {
				bp := b.path
				bp = append(bp, val.key())
				nb := bckt{bp}
				fmt.Printf("[Bucket] %s%s/\n",val.bucketString(),val.key())
				ks:=0
				bs:=0
				for _,val := range nb.getAll(){
					if val.isBucket(){
						bs++
					} else {
						ks++
					}
				}
				fmt.Printf("-- Contains %d value(s) and %d buckets.\n",ks,bs)
				if rMax<=rCur{
					fmt.Printf("-- Contents Outside Recurse Depth (Depth = %d)\n",rMax)
				}
				fmt.Println()
			} else {
				fmt.Printf("%s/%s/",rpath,val.key())
			}
			if rMax>rCur {
				bp := b.path
				bp = append(bp, val.key())
				nb := bckt{bp}
				recurseDump(nb, rMax, rCur+1, verbose)
			} else if !verbose {
				fmt.Print("...\n")
			}
		} else {
			if verbose {
				fmt.Printf("[Key] %s%s\n-- Value ([]Byte): %v\n\n",val.bucketString(),val.key(),val.v)
			} else {
				fmt.Printf("%s/%s\n",rpath,val.key())
			}
		}
	}
}
