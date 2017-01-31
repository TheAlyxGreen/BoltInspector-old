package main

import (
	"fmt"
	"strings"
)

func list(cmd []string) int{
	show := 0 // 0=all,1=buckets,2=keys
	verbose := false
	if len(cmd)>1{
		args := strings.Split(cmd[1]," ")
		for i:=0;i<len(args);i++{
			if args[i]=="b"{
				show=1
			} else if args[i]=="k" {
				show = 2
			} else if args[i]=="v" {
				verbose=true
			}
		}
	}
	bs := currentBucket.getAll()
	bs = sortArray(bs)
	bckts := []dbVal{}
	keys := []dbVal{}
	for i:=0;i<len(bs);i++{
		if bs[i].isBucket() {
			bckts=append(bckts,bs[i])
		} else {
			keys=append(keys,bs[i])
		}
	}
	if show==0||show==1 {
		for i:=0;i<len(bckts);i++{
			if !verbose {
				fmt.Println("./" + bckts[i].key())
			} else {
				fmt.Println("[Bucket] "+bckts[i].bucketString() + bckts[i].key())
				fmt.Println()
			}
		}
	}
	if show==0||show==2 {
		for i:=0;i<len(keys);i++{
			if !verbose {
				fmt.Println(keys[i].key())
			} else {
				fmt.Println("[Key] "+keys[i].bucketString() + keys[i].key())
				fmt.Print(" -- Value ([]Byte): ")
				fmt.Println(keys[i].v)
				fmt.Println()
			}
		}
	}
	return 0
}