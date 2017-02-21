# Bolt Inspector Changelog

### February 20th, 2017
```
- Added 'copy' and 'move' to the help command
- Rewrote 'help' command to have per-command help
```

### February 19th, 2017
```
- Added the 'copy' command to copy values to another key. Doesn't work with buckets (yet)
- Added the 'move' command to move value from one key to another. Uses copy command, then deletes old key.
- Changed the way strings were parsed to use shellWords ( https://github.com/mattn/go-shellwords )
- Consolidated various functions into the relevant classes to remove redundant code (especially in 'list'/'rlist')
- Reworked some minor things to make them work better 
- Updated Readme file to show links section
- Created Devlog
```

### February 15th, 2017 (part 2)
```
- Fixed 'write' and 'print' commands to correctly handle int32 and uint32
- Fixed 'write' command crashing if used in root directory
```

### February 15th, 2017 (part 1)
```
- Added the 'empty' command to delete everything in the current bucket, but keep the bucket
- Added the 'rlist' command to recursively list bucket contents
- Updated Readme file to reflect that all vital commands now exist
- Deleted 'Progress' from the Readme file, moved content to new Changelog file
```

### February 10th, 2017 (part 2)
```
- Added the 'delete' command for deleting both key/val pairs and buckets
- Added the 'bucket' command for creating a nested bucket
```

### February 10th, 2017 (part 1)
```
- Added the 'write' command to add new values to the database
- Updated the help command for 'print' (added info about arguments)
- Renamed a few files and moved some functions around
```

### February 4th, 2017
```
- Added the 'print' command to output the values of given keys
```

### January 30th, 2017 (part 2)
```
- Added the 'cd' command to move between buckets
- Updated a lot of the files, changed some underlying things around
```

### January 30th, 2017 (part 1)
```
- Initial Commit. It only has 2 functions - Help and List
```
