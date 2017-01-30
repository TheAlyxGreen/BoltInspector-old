package main

import "fmt"

func help(){
	fmt.Println(`
| ------------ BOLT INSPECTOR COMMANDS ------------ |

	All commands and arguments are case sensitive

	help:
		See this information again.

	exit:
		Terminate the program.

	list [k|b]:
		List all of the keys in the current
		bucket.
		[ARGUMENTS]
		k : only normal keys, not buckets
		b : only buckets, not normal keys


| ------------      END  COMMANDS      ------------ |
	`)
}
