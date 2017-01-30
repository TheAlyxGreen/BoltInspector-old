# Bolt Inspector
Basic database tool for [BoltDB](https://github.com/boltdb/bolt/)

## About
After working with Bolt for several days, I got incredibly fed-up with working in the dark. I searched around for a tool that would let me quickly see inside the database, but found nothing that fit my needs. So I created Bolt Inspector.

Bolt Inspector is a command-line program that quickly lets you load a database and inspect its structures and values. Unlike other programs that exist, it is built specifically to work with nested buckets.

### Features
*Features I consider necessary for the program to be considered complete*
- [x] Open Database files (a low bar, I know)
- [x] List all root buckets in database
- [x] List files in current bucket
- [x] Distinguish between keys and nested buckets when listing
- [ ] Change current bucket
- [ ] Recursive listing
- [ ] Insert new key/bucket
- [ ] Delete key/bucket
- [ ] Edit existing key
- [ ] Empty bucket (delete all keys in bucket, but not bucket)

*Features I would like to add, but may not*
- [ ] Web interface
- [ ] Rename key/bucket
- [ ] Move key/bucket

### Progress
January 30th, 2017
```
Initial Commit. It only has 2 functions - Help and List.
I'm about to begin working on the Change Bucket command
```
