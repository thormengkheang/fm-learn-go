package main

import (
	"fmt"

	"fm.com/go/files/fileutils"
)

func main() {
	filePath := "data/text.txt"
	c, err := fileutils.ReadTextFile(filePath)
	if err != nil {
		fmt.Println("ERROR")
	} else {
		fmt.Println(c)
		newContent := fmt.Sprintf("Original: %v\nDouble the original: %v%v", c, c, c)
		fileutils.WriteToFile(filePath+".output.txt", newContent)
	}
}
