package main

import (
	"fmt"
)

func help(cmd []string){

	args,r := parseArguments(cmd,0)

	if r==1 {
		fmt.Println(`
| ------------ BOLT INSPECTOR COMMANDS ------------ |

	All commands and arguments are case sensitive

	help [?]:
		See this information again, or get more
		information about a specific command.

	bucket <path>:
		Create a new bucket at the given path.

	cd <path>:
		Change the current bucket to the
		specified one.

	copy <source path> <destination path>:
		Copy the value from one path to
		another path.

	delete <path>:
		Delete the given key. Works for both
		key/value pairs and for buckets.

	empty [path]:
		Delete all values and buckets stored
		within the given bucket.

	exit:
		Terminate the program.

	list [k|b] [v]:
		List all of the keys in the current
		bucket.

	move <source path> <destination path>:
		Move a value from a given path to
		another location.

	print <path> [s|i|ui|b]:
		Print the value of the specified key.

	rlist [v] [depth=<int>]:
		Recursively list contents of the current
		bucket and all nested buckets.

	write <path> <value> [s|i|ui]:
		Write the value to the specified key.
`)
		return
	} else if r==3{
		fmt.Println("[Error] Couldn't parse arguments.")
		return
	}

	valid   := false
	constr  := args[0]
	title   := ""
	body    := ""
	argstr  := ""
	template:= `
Command:
	- %s
Title:
	- %s
Function:
	- %s
%s

`

	if args[0]=="exit" {
		valid = true
		title = "Exit"
		body = "This command terminates the program"
	} else if args[0]=="list" {
		valid = true
		title = "List"
		constr = "list [k|b] [v]"
		body = "List all of the keys in the current bucket."
		argstr = `Arguments:
	- k : only normal keys, not buckets
	- b : only buckets, not normal keys
	- v : verbose mode. Shows more information`
	} else if args[0]=="rlist" {
		valid = true
		title = "Recursive List"
		constr = "rlist [v] [depth=<int>]"
		body = `Recursively list contents of the current
	| bucket and all nested buckets.`
		argstr = `Arguments:
	- v : verbose mode. Shows more information
	- d : synonym for depth. d=<int> works.
	- depth : Depth to recurse`
	} else if args[0]=="cd" {
		valid = true
		title = "Change Bucket"
		constr = "cd <path>"
		body = `Change the current bucket to the
	| specified one. Can use absolute or
	| relative paths. Relative paths are basic
	| at the moment, not full unix emulation.
	| '~' is the database root.`
		argstr = `Arguments:
	- path: Path to the bucket being moving to`
	} else if args[0]=="print" {
		valid = true
		title = "Print"
		constr = "print <path> [s|i|ui|b]"
		body = `Print the value of the specified key.
	| You can put as many of these arguments
	| as needed, separated by spaces, and it
	| will print them all in the order you
	| provided.`
		argstr = `Arguments:
	- path: path to value to print
	- s : print as string
	- i : print as signed int32
	- ui: print as unsigned int32
	- b : print as bytes`
	} else if args[0]=="write" {
		valid = true
		title = "Write"
		constr = "write <path> <value> [s|i|ui]"
		body = `Write the value to the specified key.
	| If the key doesn't exist, it will be
	| created. If it does exist, the current
	| value will be overwritten with the new
	| value. Key is converted from string to
	| bytes to insert. Value is, by default,
	| written as the string value of the provided
	| input. Only writes once, so only the
	| first argument sticks.`
		argstr = `Arguments:
	- path: path to save value to
	- s : write value as string -> bytes
	- i : write value as int32 -> bytes
	- ui: write value as uint32 -> bytes`
	} else if args[0]=="bucket" {
		valid = true
		title = "Create Bucket"
		constr = "bucket <key>"
		body = `Create a nested bucket in the current
	| bucket with the given key. Returns an
	| error if the key is already defined.`
		argstr = `Arguments:
	- key: name of bucket to create`
	} else if args[0]=="delete" {
		valid = true
		title = "Delete"
		constr = "delete <path>"
		body = `Delete the given key. Works for both
	| key/value pairs and for buckets. It
	| asks for a confirmation prior to
	| affecting the database.`
		argstr = `Arguments:
	- path: path to value or bucket to delete`
	} else if args[0]=="empty" {
		valid = true
		title = "Empty"
		constr = "empty [path]"
		body = `Delete all values and buckets stored
	| within the given bucket. If no bucket
	| argument is given, it empties the
	| current bucket.`
		argstr = `Arguments:
	- path: path to bucket to empty`
	} else if args[0]=="move" {
		valid = true
		title = "Move"
		constr = "move <source path> <destination path>"
		body = `Move a value from a given path to another
	| location. It functions by copying the value,
	| then deleting the original. The same function
	| does both, with just a bool at the end to
	| tell it whether to delete the original or not.
	| Can also be seen as a rename command.`
		argstr = `Arguments:
	- source path: path of source value
	- destination path: path to move the value to`
	} else if args[0]=="copy" {
		valid = true
		title = "Copy"
		constr = "copy <source path> <destination path>"
		body = `Copy the value from one path to
	| another path.`
		argstr = `Arguments:
	- source path: path of source value
	- destination path: path to copy the value to`
	}

	if valid {
		fmt.Printf(template,constr,title,body,argstr)
	} else {
		fmt.Printf("[Error] Unknown Command %s\n",args[0])
	}
}