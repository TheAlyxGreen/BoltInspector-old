package main

import (
	"fmt"
	"strings"
	"encoding/binary"
	"strconv"
	"bytes"
)

func write(cmd []string){
	if len(cmd)==1{
		fmt.Println("[Error] You must specify a key and value to write.")
		return
	}

	args := strings.Split(cmd[1]," ")

	if len(args)<2{
		fmt.Println("[Error] You must specify a key and value to write.")
		return
	}

	dt:="string"

	if len(args)>2{
		if args[2]=="s"{
			dt="string"
		} else if args[2]=="i" {
			dt="int"
		}
	}

	if dt=="string"{
		currentBucket.insert([]byte(args[0]),[]byte(args[1]))
	} else if dt=="int" {
		i, err := strconv.Atoi(args[1])
		if err != nil{
			fmt.Println("[Error] Couldn't convert value string to int")
			return
		}
		buf := new(bytes.Buffer)
		err2 := binary.Write(buf, binary.LittleEndian, i)
		if err2 != nil {
			fmt.Println("binary.Write failed:", err2)
			return
		}
		currentBucket.insert([]byte(args[0]),buf.Bytes())
		return
	} else {
		fmt.Println("[Error] Unknown Insert Type.")
		return
	}
}