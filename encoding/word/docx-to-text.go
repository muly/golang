package main

import (
	"fmt"
	"os"

	"code.sajari.com/docconv"
)

func main() {
	f, err := os.Open("journal.docx")
	if err != nil {
		fmt.Println(err)
		return
	}

	text, m, err := docconv.ConvertDocx(f)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(m) // Question: what is this map for?
	fmt.Println("///////////////////////////////////////////////////////////////////////")
	fmt.Println(text)

}

//TODO: need more testing with more variety of docx files
