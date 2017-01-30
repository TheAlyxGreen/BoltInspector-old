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
// path to current bucket inside the database; "." is root
var bucketPath = []string{".","TVShows","tt2193021"}

func main() {

	// Get the database path from the user and verify it exists
	for {
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
				fmt.Println("The specified file does not exist.")
				fmt.Println("")
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
		fmt.Print("["+filename+"] "+ bpToStr(bucketPath)+" $>")
		scan := bufio.NewScanner(os.Stdin)
		scan.Scan()
		cmd := strings.SplitN(scan.Text()," ",2)

		if cmd[0]=="exit" {
			fmt.Println("Exiting...")
			break
		} else if cmd[0]=="help" {
			help()
		} else if cmd[0]=="list"{
			e:=list(cmd)
			if e==0 {
			} else if e==1{
				fmt.Println("[Error] Bucket path "+ bpToStr(bucketPath)+" is invalid. Returning to root...")
				bucketPath=[]string{"."}
			} else {
				println("Unknown Error")
			}
		} else {
			fmt.Println("Unrecognized command. Type \"help\" to see commands")
		}
	}
}