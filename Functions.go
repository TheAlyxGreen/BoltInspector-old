package main

import "os"

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
