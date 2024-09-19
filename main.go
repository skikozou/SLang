package main

import (
	"Slang/Slang/lexer"
	"Slang/Slang/parser"
	"fmt"
)

func main() {
	fmt.Println("SL test")
	lCode := lexer.Lexer("var int n < 10;var int m < 5;n <- m;")
	pCode := parser.Parser(lCode)
	fmt.Println(lCode, *pCode)

}
