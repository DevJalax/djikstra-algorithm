package main

import (
	"bytes"
	"fmt"
)

func main() {

	var strBuffer bytes.Buffer
	strBuffer.WriteString("Gideon")
	strBuffer.WriteString("Kumar")
	fmt.Println("The string buffer output is", strBuffer.String())
}
