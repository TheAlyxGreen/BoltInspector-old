# Bolt Inspector
Basic database tool for [BoltDB](https://github.com/boltdb/bolt/)

## About
After working with Bolt for several days, I got incredibly fed-up with working in the dark. I searched around for a tool that would let me quickly see inside the database, but found nothing that fit my needs. So I created Bolt Inspector.

Bolt Inspector is a command-line program that quickly lets you load a database and inspect its structures and values. Unlike other programs that exist, it is built specifically to work with nested buckets.

## Features
*Features in the most current version*
- [x] Open Database files
- [x] Move between buckets
- [x] List everything in the current bucket
- [x] Recursively list everything in the current bucket and nested buckets
- [x] Distinguish between values and nested buckets when listing
- [x] Print values (String, Int, or UInt)
- [X] Insert new values/buckets
- [x] Move/Copy existing values
- [X] Edit existing values
- [x] Delete values/buckets
- [x] Empty bucket (delete all values/buckets inside a bucket, but not the bucket itself)

*Features I would like to add, but may not*
- [ ] Web interface

## Links

For information on the program itself and how to use it, check out [the Wiki](https://github.com/89yoyos/BoltInspector/wiki)

This project is covered under the MIT License. Details can be found [here](./LICENSE)

For the latest progress, check out the [Change Log](./docs/CHANGELOG.md)

For information on future developments, check out the [Developer Log](./docs/DEVLOG.md)
