package main

import (
	"fmt"
)

func print(cmd []string){

	args,r := parseArguments(cmd,1)

	if r==2{
		fmt.Println("[Error] You must specify a key to print.")
		return
	} else if r==3{
		fmt.Println("[Error] Couldn't parse arguments.")
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
		fmt.Printf("Value for key %s in %s is undefined\n",trgt,nb.bucketString())
		return
	}

	if val.v == nil{
		fmt.Printf("Key %s is a bucket\n",trgt)
		return
	}

	if len(args)>1{
		for i:=1;i<len(args);i++{
			if args[i]=="s" || args[i]=="-s"{
				fmt.Printf("String: %s\n",val.val())
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
				fmt.Printf("Byte String: %v",val.v)
			}
		}
	} else {
		fmt.Printf("Byte String: %v",val.v)
	}

}

func readInt32(b []byte) int32 {
	// equivalent of return int32(binary.LittleEndian.Uint32(b))
	return int32(uint32(b[0]) | uint32(b[1])<<8 | uint32(b[2])<<16 | uint32(b[3])<<24)
}