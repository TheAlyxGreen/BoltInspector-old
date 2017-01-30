package main

import (
	"github.com/boltdb/bolt"
	"log"
	"fmt"
)

func list(DBPath string, bucketPath []string, cmd []string) int{
	db, err := bolt.Open(DBPath, 0600, nil)
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}

	if len(bucketPath)==1{
		db.View(func(tx *bolt.Tx) error {
			return tx.ForEach(func(name []byte, _ *bolt.Bucket) error {
				fmt.Println("./"+string(name))
				return nil
			})
		})
	} else {
		rootBuckets := []string{}
		db.View(func(tx *bolt.Tx) error {
			return tx.ForEach(func(name []byte, _ *bolt.Bucket) error {
				rootBuckets=append(rootBuckets,string(name))
				return nil
			})
		})
		exc := 0
		if contains(rootBuckets,bucketPath[1]) {
			db.View(func(tx *bolt.Tx) error {
				allBuckets := []*bolt.Bucket{}
				allBuckets = append(allBuckets,tx.Bucket([]byte(bucketPath[1])))

				for i:=2;i<len(bucketPath);i++{
					keys := []string{}
					//for each key in parent bucket
					allBuckets[i-2].ForEach(func(k,v []byte) error {
						keys=append(keys,string(k))
						return nil
					})
					if contains(keys,bucketPath[i]){
						allBuckets = append(allBuckets,allBuckets[i-2].Bucket([]byte(bucketPath[i])))
					} else {
						exc = 2
						break
					}
				}
				listbuckets := []string{}
				listkeys := []string{}
				allBuckets[len(allBuckets)-1].ForEach(func(k, v []byte) error {
					if v==nil{
						listbuckets=append(listbuckets,string(k))
					} else {
						listkeys=append(listkeys,string(k))
					}
					return nil
				})
				for i:=0;i<len(listbuckets);i++{
					fmt.Println("b | "+listbuckets[i])
				}
				fmt.Println("-")
				for i:=0;i<len(listkeys);i++{
					fmt.Println("k | "+listkeys[i])
				}

				return nil
			})
			return exc
		} else {
			return 1
		}
	}
	return 0
}