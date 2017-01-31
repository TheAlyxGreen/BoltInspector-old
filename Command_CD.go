package main

import (
	"fmt"
	"strings"
)

func cd(cmd []string){

	if len(cmd)==1{
		fmt.Println("[Error] You must specify a path.")
		return
	}

	target := cmd[1]
	bp := []string{}

	if target=="~"{
		bp = []string{"~"}
	} else if strings.HasPrefix(target,"./"){

		// if the path is relative
		bp = currentBucket.path
		target = strings.Replace(target,"./","",1)
		tmp := escapedSplit(target)
		for i:=0;i<len(tmp);i++{
			bp = append(bp,tmp[i] )
		}
	} else if strings.HasPrefix(target,"/")||strings.HasPrefix(target,"~/"){

		// if the path is absolute
		if strings.HasPrefix(target,"/"){
			target = "~"+target
		}
		tmp := escapedSplit(target)
		for i:=0;i<len(tmp);i++{
			bp = append(bp,tmp[i] )
		}
	} else {

		// assume the path is absolute
		tmp := escapedSplit(target)
		bp = append(bp,"~")
		for i:=0;i<len(tmp);i++{
			bp = append(bp,tmp[i] )
		}
	}

	nb := bckt{bp}

	if nb.exists(){
		currentBucket.path=bp
	} else {
		fmt.Println("[Error] Invalid bucket path: " + nb.bucketString())
	}

}

func escapedSplit(s string) []string{
	if !strings.Contains(s,"\\/"){
		return strings.Split(s,"/")
	}
	rndm := RandString(10)
	for i:=0;strings.Contains(s,rndm);i++{
		rndm = RandString(10)
	}
	s = strings.Replace(s,"\\/",rndm,-1)
	tmp := strings.Split(s,"/")
	for i:=0;i<len(tmp);i++{
		tmp[i]=strings.Replace(tmp[i],rndm,"/",-1)
	}
	return tmp
}