package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)


func main()  {
	
	fmt.Println("Welcome to Word Frequency Counter")

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter File Path (with Name and Extension)")
	filePath, _ := reader.ReadString('\n')
	filePath = strings.TrimSpace(filePath)

	dataByte, err := os.ReadFile(filePath)
	checkErr(err)

	fileString := string(dataByte)

	fields := strings.Fields(strings.ToLower(fileString))

	stringCountMap := make(map[string]int)

	for _, word := range fields {
		stringCountMap[word]++;
	}

	var newFileContent string

	for word, count := range stringCountMap {
		printLine := fmt.Sprintf("Word '%v' is used %v times in this file", word, count)
		newFileContent += printLine+"\n" 
	}

	fmt.Println(newFileContent)

	newFile, _ := os.Create("output.txt")
	defer newFile.Close()

	io.WriteString(newFile, newFileContent)
}


func checkErr (err error) {
	if err != nil {
		panic(err)
	}
}