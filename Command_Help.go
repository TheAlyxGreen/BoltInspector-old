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
		You can put as many of these arguments
		as needed, separated by spaces, and it
		will print them all in the order you
		provided.
		[ARGUMENTS]
		s : print as string
		i : print as signed int
		ui: print as unsigned int
		b : print as bytes

	write <key> <value> [s|i]:
		Write the value to the specified key.
		If the key doesn't exist, it will be
		created. If it does exist, the current
		value will be overwritten with the new
		value. Key is converted from string to
		bytes to insert. Value is, by default,
		written as ascii values of provided
		string. Only writes once, so only the
		first argument sticks.
		[ARGUMENTS]
		s : write value as string -> bytes
		i : write value as int -> bytes

	bucket <key>:
		Create a nested bucket in the current
		bucket with the given key. Returns an
		error if the key is already defined.

	delete <key>:
		Delete the given key. Works for both
		key/value pairs and for buckets. It
		asks for a confirmation prior to
		affecting the database.


| ------------      END  COMMANDS      ------------ |
	`)
	}

	// Adding command-specific help in the future...
	// Will move arguments into command-specific help
}