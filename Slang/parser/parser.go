package parser

import (
	"Slang/Slang"
	"errors"
	"strconv"
)

func Parser(s [][]string) *Slang.Code {
	parsedCode := Slang.Code{
		Elem: make(Slang.ElemMap),
	}
	for _, l := range s {
		switch l[0] {
		case "var":
			if len(l) >= 5 && l[3] == "<" {
				switch l[1] {
				case "int":
					var err error
					parsedCode.Elem[l[2]], err = strconv.Atoi(l[4])
					if err != nil {
						Slang.Err(err)
					}
				case "string":
					parsedCode.Elem[l[2]] = l[4]
				case "bool":
					if l[4] == "true" {
						parsedCode.Elem[l[2]] = true
					} else if l[4] == "false" {
						parsedCode.Elem[l[2]] = false
					} else {
						Slang.Err(errors.New("Invalid value"))
					}
				default:
					continue
				}
			}
		case "Printf":

		default:
			if val, ok := parsedCode.Elem[l[0]]; ok {
				switch val.(type) {
				case int:
					switch l[1] {
					case "<":
						if val, ok := parsedCode.Elem[l[2]]; ok {
							switch val.(type) {
							case int:
								parsedCode.Elem[l[0]] = parsedCode.Elem[l[2]]
							}
						}
					case "<+":
						if val, ok := parsedCode.Elem[l[2]]; ok {
							switch val.(type) {
							case int:
								parsedCode.Elem[l[0]] = parsedCode.Elem[l[0]].(int) + parsedCode.Elem[l[2]].(int)
							}
						}
					case "<-":
						if val, ok := parsedCode.Elem[l[2]]; ok {
							switch val.(type) {
							case int:
								parsedCode.Elem[l[0]] = parsedCode.Elem[l[0]].(int) - parsedCode.Elem[l[2]].(int)
							}
						}
					case "<*":
						if val, ok := parsedCode.Elem[l[2]]; ok {
							switch val.(type) {
							case int:
								parsedCode.Elem[l[0]] = parsedCode.Elem[l[0]].(int) * parsedCode.Elem[l[2]].(int)
							}
						}
					case "</":
						if val, ok := parsedCode.Elem[l[2]]; ok {
							switch val.(type) {
							case int:
								parsedCode.Elem[l[0]] = parsedCode.Elem[l[0]].(int) / parsedCode.Elem[l[2]].(int)
							}
						}
					}
				case string:

				case bool:

				}
			}
		}
	}
	return &parsedCode
}