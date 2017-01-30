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

func main() {

	path,filename := "", ""

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

	db, err := bolt.Open(path, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	db.Close()

	_, filename = filepath.Split(path)

	bucketPath := []string{".","TVShows","tt2193021"}
	for { // loop for each command
		fmt.Print("["+filename+"] "+str(bucketPath)+" $>")
		scan := bufio.NewScanner(os.Stdin)
		scan.Scan()
		cmd := strings.SplitN(scan.Text()," ",2)

		if cmd[0]=="exit" {
			fmt.Println("Exiting...")
			break
		} else if cmd[0]=="help" {
			help()
		} else if cmd[0]=="list"{
			e:=list(path,bucketPath,cmd)
			if e==0 {
			} else if e==1{
				fmt.Println("[Error] Bucket path "+str(bucketPath)+" is invalid. Returning to root...")
				bucketPath=[]string{"."}
			} else {
				println("Unknown Error")
			}
		} else {
			fmt.Println("Unrecognized command. Type \"help\" to see commands")
		}
	}
}

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil { return true, nil }
	if os.IsNotExist(err) { return false, nil }
	return true, err
}

func str(s []string) string {
	r := s[0]
	for i:=1;i<len(s);i++{
		r=r+"/"+s[i]
	}
	return "("+r+")"
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}