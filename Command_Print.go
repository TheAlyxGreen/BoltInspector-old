package main

import (
	"fmt"
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
				fmt.Print("Signed Int: ")
				if len(val.v)!=8{
					fmt.Println("Cannot be read as a Signed Int (32)")
				} else {
					data := readInt32(val.v)
					fmt.Println(data)
				}
			} else if args[i]=="ui" || args[i]=="-ui" {
				fmt.Print("Unsigned Int: ")
				if len(val.v)!=8{
					fmt.Println("Cannot be read as a Signed Int (32)")
				} else {
					data := uint32(readInt32(val.v))
					fmt.Println(data)
				}
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

func readInt32(b []byte) int32 {
	// equivalnt of return int32(binary.LittleEndian.Uint32(b))
	return int32(uint32(b[0]) | uint32(b[1])<<8 | uint32(b[2])<<16 | uint32(b[3])<<24)
}