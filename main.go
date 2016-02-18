package main

import "fmt"
import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func main() {
	fmt.Print(cases.Upper(language.English).String("Hello world"))
}
