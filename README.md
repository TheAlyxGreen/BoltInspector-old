# Bolt Inspector
Basic database tool for [BoltDB](https://github.com/boltdb/bolt/)

## About
After working with Bolt for several days, I got incredibly fed-up with working in the dark. I searched around for a tool that would let me quickly see inside the database, but found nothing that fit my needs. So I created Bolt Inspector.

Bolt Inspector is a command-line program that quickly lets you load a database and inspect its structures and values. Unlike other programs that exist, it is built specifically to work with nested buckets.

To view the latest progress, check out the [Change Log](./CHANGELOG.md)

## Features
*Features in the most current version*
- [x] Open Database files (a low bar, I know)
- [x] List all root buckets in database
- [x] List keys in current bucket
- [x] Distinguish between keys and nested buckets when listing
- [x] Change current bucket
- [x] Print key values
- [X] Insert new key/bucket
- [X] Edit existing key
- [x] Delete key/bucket
- [x] Empty bucket (delete all keys in bucket, but not bucket)
- [x] Recursive listing

*Features I would like to add, but may not*
- [ ] Web interface
- [ ] Rename key/bucket
- [ ] Move key/bucket

