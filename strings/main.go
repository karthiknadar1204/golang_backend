package main

import (
	"fmt"
	"strings"
)

func main(){
	var str string;
	str="Hello, World!"
	fmt.Println(str)

	data:="apple.oranges.banana.pears"
	parts:= strings.Split(data,".")
	fmt.Println(parts)

	joined:= strings.Join(parts,"-")
	fmt.Println(joined)

	contains:= strings.Contains(data,"apple")
	fmt.Println(contains)

	count:= strings.Count(data,"a")
	fmt.Println(count)

	replaced:= strings.Replace(data,"a","A",-1)
	fmt.Println(replaced)

	trimmed:= strings.Trim(data,".")
	fmt.Println(trimmed)

	trimmedLeft:= strings.TrimLeft(data,".")
	fmt.Println(trimmedLeft)

	trimmedRight:= strings.TrimRight(data,".")
	fmt.Println(trimmedRight)

	trimmedSpace:= strings.TrimSpace(data)
	fmt.Println(trimmedSpace)
}
