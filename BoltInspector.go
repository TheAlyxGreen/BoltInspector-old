package main

import (
	"bufio"
	"os"
	"fmt"
	"github.com/boltdb/bolt"
	"log"
	"path/filepath"
	"strings"
)

// full path to database; just the file name
var path,filename = "", ""
// var for root bucket
var root = bckt{[]string{"~"}}
// path to current bucket inside the database; "~" is root
var currentBucket = bckt{[]string{"~"}}

func main() {

	// Get the database path from the user and verify it exists
	for {
		fmt.Println()
		fmt.Print("Database to Read (or exit): ")
		scan := bufio.NewScanner(os.Stdin)
		scan.Scan()
		DBPath := scan.Text()

		ex,_ := exists(DBPath)

		if DBPath=="exit" {
			fmt.Println("Exiting...")
			os.Exit(0)
			return
		} else if !ex {
				fmt.Print("[Error] The specified file does not exist.\n")
		} else {
			path = DBPath
			break
		}
	}

	// verify that bolt can open the database
	db, err := bolt.Open(path, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	db.Close()

	// set the filename var
	_, filename = filepath.Split(path)

	// main loop of the script
	for {
		fmt.Print("["+filename+"] ("+currentBucket.bucketString()+") $>")
		scan := bufio.NewScanner(os.Stdin)
		scan.Scan()
		cmd := strings.SplitN(scan.Text()," ",2)

		if cmd[0]=="exit" {
			fmt.Println("Exiting...")
			break
		} else if cmd[0]=="help" {
			help(cmd)
		} else if cmd[0]=="list" {
			list(cmd)
		} else if cmd[0]=="rlist"{
			rlist(cmd)
		} else if cmd[0]=="cd"{
			cd(cmd)
		} else if cmd[0]=="print"{
			print(cmd)
		} else if cmd[0]=="write"{
			write(cmd)
		} else if cmd[0]=="bucket"{
			bucket(cmd)
		} else if cmd[0]=="delete"{
			delete(cmd)
		} else if cmd[0]=="empty"{
			emptyBucket(cmd)
		} else if cmd[0]=="copy"{
			copy(cmd,false)
		} else if cmd[0]=="move"{
			copy(cmd,true)
		} else {
			fmt.Println("Unrecognized command. Type \"help\" to see commands")
		}
	}
}