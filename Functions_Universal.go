package main

import (
	"os"
	"math/rand"
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

// Creates the nice string for the bucket path
func bpToStr(s []string) string {
	r := s[0]
	for i:=1;i<len(s);i++{
		r=r+"/"+s[i]
	}
	return "("+r+")"
}

// Does array contain given string?
func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func cpyBytes(s []byte) []byte{
	r := []byte{}
	for i:=0;i<len(s);i++ {
		r = append(r,s[i])
	}
	return r
}

func RandString(n int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}