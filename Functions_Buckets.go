package main

import (
	"github.com/boltdb/bolt"
	"log"
	"fmt"
)

/*

This is essentially the Bucket class and its functions
They're only separated to make the files smaller and reduce clutter

*/

// struct for bucket info and methods
type bckt struct {
	path []string
}

// return the path as a single string for printing
func (b bckt) bucketString() string{
	s := ""
	for i:=0;i<len(b.path);i++{
		s=s+ b.path[i]+"/"
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
			db.View(func(tx *bolt.Tx) error {

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

// return an individual result from the bucket, as well as whether the requested key exists
func (b bckt) getOne(key string) (dbVal,bool){
	vals := b.getAll()
	for i:=0;i<len(vals);i++{
		if string(vals[i].k) == key{
			return vals[i], true
		}
	}
	return dbVal{}, false
}

// return just the values, not the buckets
func (b bckt) getValues() []dbVal{
	vals := []dbVal{}
	for _,val := range b.getAll(){
		if !val.isBucket() {
			vals = append(vals, val)
		}
	}
	return vals
}

// return just the buckets, not the values
func (b bckt) getBuckets() []dbVal{
	vals := []dbVal{}
	for _,val := range b.getAll(){
		if val.isBucket() {
			vals = append(vals, val)
		}
	}
	return vals
}

// returns 2 separate arrays, one for buckets and one for values. More efficient if getting both
func (b bckt) getAllSeparated() ([]dbVal,[]dbVal){
	bs := []dbVal{}
	vs := []dbVal{}
	for _,val := range b.getAll(){
		if val.isBucket() {
			bs = append(bs, val)
		} else {
			vs = append(vs, val)
		}
	}
	return bs,vs
}

// return the total number of entries in the bucket
func (b bckt) count() int{
	return len(b.getAll())
}

// return the number of non-bucket values stored in this bucket
func (b bckt) valueCount() int {
	count := 0
	for _,val := range b.getAll(){
		if !val.isBucket(){
			count++
		}
	}
	return count
}

// return the number of buckets nested inside this bucket
func (b bckt) bucketCount() int {
	count := 0
	for _,val := range b.getAll(){
		if val.isBucket(){
			count++
		}
	}
	return count
}

func (b bckt) countBoth() (int,int){
	bc := 0
	vc := 0
	for _,val := range b.getAll(){
		if val.isBucket(){
			bc++
		} else {
			vc++
		}
	}
	return bc,vc
}

// insert the specified value at the specified key
func (b bckt) insert(key []byte, val []byte) bool{
	if !b.exists() {
		return false
	}

	db, err := bolt.Open(path, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.Update(func(tx *bolt.Tx) error {
		return tx.ForEach(func(name []byte, _ *bolt.Bucket) error {
			// create array to store references to buckets
			allBuckets := []*bolt.Bucket{}

			// set first bucket to root bucket
			allBuckets = append(allBuckets,tx.Bucket([]byte(b.path[1])))

			// burrow into bottom bucket of path
			for i:=2;i<len(b.path);i++{
				allBuckets = append(allBuckets,allBuckets[i-2].Bucket([]byte(b.path[i])))
			}

			allBuckets[len(allBuckets)-1].Put(key, val)

			return nil
		})
	})

	return true
}

// create a nested bucket with the specified key
func (b bckt) insertBucket(key []byte) bool {
	if !b.exists() {
		return false
	}

	db, err := bolt.Open(path, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.Update(func(tx *bolt.Tx) error {
		if len(b.path)>1 {
			return tx.ForEach(func(name []byte, _ *bolt.Bucket) error {
				// create array to store references to buckets
				allBuckets := []*bolt.Bucket{}

				// set first bucket to root bucket
				allBuckets = append(allBuckets, tx.Bucket([]byte(b.path[1])))

				// burrow into bottom bucket of path
				for i := 2; i < len(b.path); i++ {
					allBuckets = append(allBuckets, allBuckets[i-2].Bucket([]byte(b.path[i])))
				}

				allBuckets[len(allBuckets)-1].CreateBucketIfNotExists(key)

				return nil
			})
		} else {
			_,err := tx.CreateBucket(key)
			if err!=nil {
				return fmt.Errorf("[Error] Failed to create bucket: %s", err)
			}
			return nil
		}
	})

	return true
}

// delete the value or bucket at the specified key
func (b bckt) delete(key []byte) bool{

	if !b.exists() {
		return false
	}

	db, err := bolt.Open(path, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.Update(func(tx *bolt.Tx) error {
		if len(b.path)>1 {
			return tx.ForEach(func(name []byte, _ *bolt.Bucket) error {
				// create array to store references to buckets
				allBuckets := []*bolt.Bucket{}

				// set first bucket to root bucket
				allBuckets = append(allBuckets, tx.Bucket([]byte(b.path[1])))

				// burrow into bottom bucket of path
				for i := 2; i < len(b.path); i++ {
					allBuckets = append(allBuckets, allBuckets[i-2].Bucket([]byte(b.path[i])))
				}

				if allBuckets[len(allBuckets)-1].Get(key) != nil {
					allBuckets[len(allBuckets)-1].Delete(key)
				} else {
					allBuckets[len(allBuckets)-1].DeleteBucket(key)
				}

				return nil
			})
		} else {
			tx.DeleteBucket(key)
			return nil
		}
	})

	return true
}

// delete all values and buckets located within this bucket
func (b bckt) empty() bool{

	if !b.exists() {
		return false
	}

	db, err := bolt.Open(path, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.Update(func(tx *bolt.Tx) error {
		return tx.ForEach(func(name []byte, _ *bolt.Bucket) error {
			// create array to store references to buckets
			allBuckets := []*bolt.Bucket{}

			// set first bucket to root bucket
			allBuckets = append(allBuckets,tx.Bucket([]byte(b.path[1])))

			// burrow into bottom bucket of path
			for i:=2;i<len(b.path);i++{
				allBuckets = append(allBuckets,allBuckets[i-2].Bucket([]byte(b.path[i])))
			}

			allBuckets[len(allBuckets)-1].ForEach(func(k, v []byte) error {
				if v!=nil{
					allBuckets[len(allBuckets)-1].Delete(k)
				} else {
					allBuckets[len(allBuckets)-1].DeleteBucket(k)
				}
				return nil
			})

			return nil
		})
	})

	return true
}

func (b bckt) isRoot() bool{
	if len(b.path)==1{
		if b.path[0]=="~" {
			return true
		}
	}
	return false
}