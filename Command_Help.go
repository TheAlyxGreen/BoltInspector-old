package main

import "fmt"

func help(cmd []string){
	if len(cmd)==1 {
		fmt.Println(`
| ------------ BOLT INSPECTOR COMMANDS ------------ |

	All commands and arguments are case sensitive

	help [?]:
		See this information again.
		[ARGUMENTS]
		? : type "help" plus the name of another
		    command to get help with that
		    specific command

	exit:
		Terminate the program.

	list [k|b] [v]:
		List all of the keys in the current
		bucket.
		[ARGUMENTS]
		k : only normal keys, not buckets
		b : only buckets, not normal keys
		v : verbose mode. Shows more information

	cd <path>:
		Change the current bucket to the
		specified one. Can use absolute or
		relative paths. Relative paths are basic
		at the moment, not full unix emulation.
		~ is the database root.

	print <key> [s|i|ui|b]:
		Print the value of the specified key.
		For now, the key must be in the bucket
		currently being inspected. If you need
		to read a key in a different bucket,
		you need to CD there first.


| ------------      END  COMMANDS      ------------ |
	`)
	}

	// Adding command-specific help in the future...
	// Will move arguments into command-specific help
}