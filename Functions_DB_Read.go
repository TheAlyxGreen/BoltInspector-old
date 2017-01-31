package main

import (
	"github.com/boltdb/bolt"
	"log"
)

/*

This file contains functions related to reading the database.
They're only separated to make the files smaller and reduce clutter

*/

// struct for bucket info and methods
type bckt struct {
	path []string
}

func (d bckt) bucketString() string{
	s := ""
	for i:=0;i<len(d.path);i++{
		s=s+d.path[i]+"/"
	}
	return s
}

// reset the bucket to the root
func (b *bckt) reset(){
	b.path = []string{"~"}
}

// verify this bucket exists
func (b bckt) exists() bool{
	bp := b.path
	if len(bp)==1{
		if bp[0]=="~"{
			return true
		} else {
			return false
		}
	}
	bucketExists := true
	db, err := bolt.Open(path, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	rootBuckets := []string{}
	db.View(func(tx *bolt.Tx) error {
		return tx.ForEach(func(name []byte, _ *bolt.Bucket) error {
			rootBuckets=append(rootBuckets,string(name))
			return nil
		})
	})
	if contains(rootBuckets,bp[1]) {
		if len(bp) > 2 {
			db.View(func(tx *bolt.Tx) error {
				allBuckets := []*bolt.Bucket{}
				allBuckets = append(allBuckets, tx.Bucket([]byte(bp[1])))
				for i := 2; i < len(bp); i++ {
					keys := []string{}
					//for each key in parent bucket
					allBuckets[i - 2].ForEach(func(k, v []byte) error {
						keys = append(keys, string(k))
						return nil
					})
					if contains(keys, bp[i]) {
						allBuckets = append(allBuckets, allBuckets[i - 2].Bucket([]byte(bp[i])))
					} else {
						bucketExists = false
						break
					}
				}
				return nil
			})
		}
	} else {
		bucketExists = false
	}
	return bucketExists
}

// return an array of all key/value pairs in the bucket
func (b bckt) getAll() []dbVal{
	// create array we'll return later
	r := []dbVal{}

	// verify that specified bucket path exists
	if b.exists(){

		// open db for reading
		db, err := bolt.Open(path, 0600, nil)
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		if len(b.path)==1{

			// if path is root, dump root buckets
			db.View(func(tx *bolt.Tx) error {
				return tx.ForEach(func(name []byte, _ *bolt.Bucket) error {
					t := dbVal{}
					t.path = b.path
					t.v = nil
					t.k=cpyBytes(name)
					r = append(r,t)
					return nil
				})
			})
		} else {
			// if path isn't the root, recurse into path
			db.Update(func(tx *bolt.Tx) error {

				// create array to store references to buckets
				allBuckets := []*bolt.Bucket{}

				// set first bucket to root bucket
				allBuckets = append(allBuckets,tx.Bucket([]byte(b.path[1])))

				// burrow into bottom bucket of path
				for i:=2;i<len(b.path);i++{
					allBuckets = append(allBuckets,allBuckets[i-2].Bucket([]byte(b.path[i])))
				}

				// for all in last bucket, copy values to r
				allBuckets[len(allBuckets)-1].ForEach(func(k, v []byte) error {
					t := dbVal{}
					t.path = b.path
					t.k=cpyBytes(k)
					if v!=nil {
						t.v = cpyBytes(v)
					} else {
						t.v = nil
					}
					r = append(r,t)
					return nil
				})
				return nil
			})
		}
	} else {
		return nil
	}

	// return the results
	return r
}