package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main()  {

	type Student struct {
		Name string
		RollNo int
		Marks [3]int
	}

	students := []Student{{
		Name: "Alice",
		RollNo: 1,
		Marks: [3]int{85, 90, 95},
	}, {
		Name: "Bob",
		RollNo: 2,
		Marks: [3]int{80, 85, 90},
	}, {
		Name: "Charlie",
		RollNo: 3,
		Marks: [3]int{75, 80, 85},
	}}

	rollNoMap := make(map[int]*Student);

	rollNoMap[1] = &students[0]
	rollNoMap[2] = &students[1]
	rollNoMap[3] = &students[2]


	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter Roll No:");
	rollNo, _ := reader.ReadString('\n')

	rollNoInInt, _ := strconv.Atoi(strings.TrimSpace(rollNo));

	address := rollNoMap[rollNoInInt]
	fetchedStudent := *address

	fmt.Println("Name:", fetchedStudent.Name)
	fmt.Println("Marks:", fetchedStudent.Marks)
	fmt.Println("Average Marks:", (fetchedStudent.Marks[0] + fetchedStudent.Marks[1] + fetchedStudent.Marks[2]) / 3)
	fmt.Printf("Address: %p \n", address)


	
}