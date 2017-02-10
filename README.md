# Bolt Inspector
Basic database tool for [BoltDB](https://github.com/boltdb/bolt/)

## About
After working with Bolt for several days, I got incredibly fed-up with working in the dark. I searched around for a tool that would let me quickly see inside the database, but found nothing that fit my needs. So I created Bolt Inspector.

Bolt Inspector is a command-line program that quickly lets you load a database and inspect its structures and values. Unlike other programs that exist, it is built specifically to work with nested buckets.

### Features
*Features I consider necessary for the program to be fully functioning*
- [x] Open Database files (a low bar, I know)
- [x] List all root buckets in database
- [x] List keys in current bucket
- [x] Distinguish between keys and nested buckets when listing
- [x] Change current bucket
- [x] Print key values
- [X] Insert new key/bucket
- [X] Edit existing key
- [x] Delete key/bucket
- [ ] Empty bucket (delete all keys in bucket, but not bucket)
- [ ] Recursive listing

*Features I would like to add, but may not*
- [ ] Web interface
- [ ] Rename key/bucket
- [ ] Move key/bucket

### Progress
February 10th, 2017 (part 2)
```
Added the 'delete' command for deleting both key/val
pairs and buckets, and added the 'bucket' command for
adding a nested bucket to the current bucket. Added 
the relevant help information for both
```

February 10th, 2017 (part 1)
```
Added the 'write' command, renamed a few files and
moved some functions around. I also updated the help
command for both 'print' (added info about arguments)
and for 'write'
```

February 4th, 2017
```
Added the 'print' command to output the values of given
keys. Also added the relevant section to the 'help'
command and (obviously) updated the readme file.
```

January 30th, 2017 (part 2)
```
Updated a lot of the files, changed some things around.
Added the CD (change directory) command to move around
inside of the database.
```

January 30th, 2017 (part 1)
```
Initial Commit. It only has 2 functions - Help and List.
I'm about to begin working on the Change Bucket command
```
