package main

import (
	"fmt"
	"encoding/binary"
	"strings"
)

func print(cmd []string){

	if len(cmd)==1{
		fmt.Println("[Error] You must specify a key to print.")
		return
	}

	args := strings.Split(cmd[1]," ")

	val,suc := currentBucket.getOne(args[0])

	if !suc {
		fmt.Println("Value for key \"" + args[0] + "\" is undefined in this bucket (" + currentBucket.bucketString() + ")")
		return
	}

	if val.v == nil{
		fmt.Println("Key \"" + args[0] + "\" is a bucket")
		return
	}

	if len(args)>1{
		for i:=1;i<len(args);i++{
			if args[i]=="s" || args[i]=="-s"{
				fmt.Print("String: ")
				fmt.Println(val.val())
			} else if args[i]=="i" || args[i]=="-i" {
				data,_ := binary.Varint(val.v)
				fmt.Print("Signed Int: ")
				fmt.Println(data)
			} else if args[i]=="ui" || args[i]=="-ui" {
				data,_ := binary.Uvarint(val.v)
				fmt.Print("Unsigned Int: ")
				fmt.Println(data)
			} else if args[i]=="b" || args[i]=="-b"{
				fmt.Print("Byte String: ")
				fmt.Println(val.v)
			}
		}
	} else {
		fmt.Print("Byte String: ")
		fmt.Println(val.v)
	}

}