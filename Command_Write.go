package main

import (
	"fmt"
	"encoding/binary"
	"strconv"
	"bytes"
)

func write(cmd []string){

	if len(currentBucket.path)==1{
		fmt.Println("[Error] You cannot write values in the root directory. Use the 'bucket' command to make a value storing bucket.")
		return
	}

	args,r := parseArguments(cmd,2)

	if r==2{
		fmt.Println("[Error] You must specify a key and value to write.")
		return
	} else if r==3{
		fmt.Println("[Error] Couldn't parse arguments")
		return
	}

	path := stringToPath(args[0],currentBucket)
	trgt := path[len(path)-1]

	nb := bckt{path[:len(path)-1]}

	if !nb.exists() {
		fmt.Println("[Error] The specified bucket doesn't exist")
		return
	}

	dt:="string"

	if len(args)>2{
		if args[2]=="s"{
			dt="string"
		} else if args[2]=="i" {
			dt="int"
		} else if args[2]=="ui"{
			dt="uint"
		}
	}

	if dt=="string"{
		nb.insert([]byte(trgt),[]byte(args[1]))
	} else if dt=="int" {
		i, err := strconv.ParseInt(args[1], 10, 32)
		if err != nil {
			fmt.Printf("[Error] Couldn't convert value string to int:\n%s\n", err)
			return
		}
		buf := new(bytes.Buffer)
		err2 := binary.Write(buf, binary.LittleEndian, i)
		if err2 != nil {
			fmt.Println("binary.Write failed:", err2)
			return
		}
		nb.insert([]byte(trgt), buf.Bytes())
		return
	} else if dt=="uint" {
		i, err := strconv.ParseUint(args[1], 10, 32)
		if err != nil {
			fmt.Printf("[Error] Couldn't convert value string to int:\n%s", err)
			return
		}
		buf := new(bytes.Buffer)
		err2 := binary.Write(buf, binary.LittleEndian, i)
		if err2 != nil {
			fmt.Println("[Error] binary.Write failed:", err2)
			return
		}
		nb.insert([]byte(trgt), buf.Bytes())
		return
	} else {
		fmt.Println("[Error] Unknown Insert Type.")
		return
	}
}