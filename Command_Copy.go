package main

import (
	"fmt"
)

func copy(cmd []string, del bool){

	args,r := parseArguments(cmd,2)

	if r==2{
		fmt.Println("[Error] You must enter both a target and destination key")
		return
	} else if r==3{
		fmt.Println("[Error] Couldn't parse arguments")
		return
	}

	path1 := stringToPath(args[0],currentBucket)
	path2 := stringToPath(args[1],currentBucket)

	k1 := path1[len(path1)-1]
	k2 := path2[len(path2)-1]


	if len(path1)<2 || len(path2)<2{
		fmt.Println("[Error] Invalid path(s) entered [length error]")
		return
	}

	pb1 := bckt{path1[:len(path1)-1]}
	pb2 := bckt{path2[:len(path2)-1]}

	if !pb1.exists() {
		fmt.Printf("[Error] Source bucket (%s) doesn't exist\n",path1)
		return
	}

	if !pb2.exists() {
		fmt.Printf("[Error] Destination bucket (%s) doesn't exist\n",path2)
	}

	_,e2 := pb2.getOne(k2)

	if e2{
		fmt.Println("[Error] Destination key already exists")
		return
	}

	v1,e1 := pb1.getOne(k1)

	if !e1{
		fmt.Println("[Error] Source key doesn't exist")
		return
	}

	if v1.isBucket(){
		fmt.Println("[Error] Source is a bucket. Copying buckets is not yet supported")
		return
	} else {
		pb2.insert([]byte(k2),v1.v)
	}

	if del{
		pb1.delete(v1.k)
	}

}