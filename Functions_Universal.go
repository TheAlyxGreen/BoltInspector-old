package main

import (
	"os"
	"math/rand"
	"strings"
	"github.com/mattn/go-shellwords"
)

/*

This is a function library for universally relevant functions.
Functions specific to certain commands will be in the files associated with those commands.

*/

// Check if file exists at a given path, and that the path is valid
func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil { return true, nil }
	if os.IsNotExist(err) { return false, nil }
	return true, err
}

// Does array contain given string?
func contains(haystack []string, needle string) bool {
	for _, str := range haystack {
		if str == needle {
			return true
		}
	}
	return false
}

// copies an array of bytes to prevent it disappearing when the db is closed
func cpyBytes(s []byte) []byte{
	r := []byte{}
	for i:=0;i<len(s);i++ {
		r = append(r,s[i])
	}
	return r
}

// generate a random string of n length
func RandString(n int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

// parse a string and return a bucket path. 'relativeTo' is the bucket the path is assumed to be relative to
func stringToPath(target string, relativeTo bckt) []string{
	bp := []string{}
	if target=="~"{
		bp = []string{"~"}
	} else if strings.HasPrefix(target,"./"){

		// if the path is relative
		bp = currentBucket.path
		target = strings.Replace(target,"./","",1)
		tmp := escapedSplit(target)
		for i:=0;i<len(tmp);i++{
			bp = append(bp,tmp[i] )
		}
	} else if strings.HasPrefix(target,"/")||strings.HasPrefix(target,"~/"){

		// if the path is absolute
		if strings.HasPrefix(target,"/"){
			target = "~"+target
		}
		tmp := escapedSplit(target)
		for i:=0;i<len(tmp);i++{
			bp = append(bp,tmp[i] )
		}
	} else {
		if relativeTo.exists(){
			// assume the path is relative
			bp = relativeTo.path
			tmp := escapedSplit(target)
			for i:=0;i<len(tmp);i++{
				bp = append(bp,tmp[i] )
			}
		} else {
			// assume the path is absolute
			tmp := escapedSplit(target)
			bp = append(bp,"~")
			for i:=0;i<len(tmp);i++{
				bp = append(bp,tmp[i] )
			}
		}
	}
	return bp
}

// split a string for bucket path, but allow escaping backslashes
func escapedSplit(s string) []string{
	if !strings.Contains(s,"\\/"){
		return strings.Split(s,"/")
	}
	rndm := RandString(10)
	for i:=0;strings.Contains(s,rndm);i++{
		rndm = RandString(10)
	}
	s = strings.Replace(s,"\\/",rndm,-1)
	tmp := strings.Split(s,"/")
	for i:=0;i<len(tmp);i++{
		tmp[i]=strings.Replace(tmp[i],rndm,"/",-1)
	}
	return tmp
}

func plurality(singular string, plural string, count int) string{
	if count==1{
		return singular
	}
	return plural
}

func valPlural(count int) string{
	return plurality("value","values",count)
}

func bcktPlural(count int) string{
	return plurality("bucket","buckets",count)
}

func parseArguments(cmd []string, minArgs int) (args []string, returnCode int){

	hasArgs := len(cmd)>1

	if !hasArgs && minArgs<=0{
		return []string{},1 // No arguments
	} else if !hasArgs && minArgs>0{
		return []string{},2 // Not enough arguments
	}

	args,err := shellwords.Parse(cmd[1])

	if err!=nil{
		return []string{},3 // Failed to parse
	}

	if len(args)<minArgs {
		return []string{},2 // Not enough arguments
	}

	return args,0
}