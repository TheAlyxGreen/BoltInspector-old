package main

import (
	"os"
	"sort"
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

// Key/Value pairs in the DB
type dbVal struct {
	path []string
	k []byte
	v []byte
}

// Get the key as a string
func (d dbVal) key() string{
	return string(d.k)
}

// Get the value as a string
func (d dbVal) val() string{
	return string(d.v)
}

// Test if it is a bucket
func (d dbVal) isBucket() bool{
	return d.v==nil
}

func (d dbVal) bucketString() string{
	s := ""
	for i:=0;i<len(d.path);i++{
		s=s+d.path[i]+"/"
	}
	return s
}

func sortArray(dbvs []dbVal)[]dbVal{
	r := []dbVal{}
	tmp := make(map[string]dbVal)
	keys := []string{}
	for i:=0;i<len(dbvs);i++{
		keys=append(keys,dbvs[i].key())
		tmp[dbvs[i].key()]=dbvs[i]
	}
	sort.Strings(keys)
	for i:=0;i<len(keys);i++{
		r=append(r,tmp[keys[i]])
	}
	return r
}