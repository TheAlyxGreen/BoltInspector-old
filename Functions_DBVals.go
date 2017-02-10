package main

import "sort"

/*

This is essentially the DBVals class and its functions
They're only separated to make the files smaller and reduce clutter

*/

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