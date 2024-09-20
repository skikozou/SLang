package main

import (
	"Slang/Slang/lexer"
	"Slang/Slang/parser"
	"fmt"
	"os"
)

func main() {
	fmt.Println("SL test")
	bin, _ := os.ReadFile("test.sl")
	lCode := lexer.Lexer(string(bin))
	pCode := parser.Parser(lCode)
	fmt.Println(lCode, *pCode)

}
