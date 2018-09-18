package main

import (
	"fmt"

	"github.com/vanhtuan0409/go-func-parser"
)

func main() {
	fSignature, err := parser.Parse("function void hello( int arg1,  string arg2)")
	if err != nil {
		panic(err)
	}
	printOutput(fSignature)
}

func printOutput(fSignature *parser.FuncSignature) {
	fmt.Printf("Function name: %s\n", fSignature.Name)
	fmt.Printf("Function return type: %s\n", fSignature.ReturnType)
	for index, param := range fSignature.ParamList {
		fmt.Printf("Param %d's name: %s\n", index+1, param.Name)
		fmt.Printf("Param %d's type: %s\n", index+1, param.Type)
	}
}
