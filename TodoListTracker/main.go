package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Task struct {
	Id int
	Task string
	IsDone bool
}

func main()  {



	tasks := []Task{}
	idTaskMap := make(map[int]*Task)
	defer printSummary(idTaskMap)

	for {
		choice := menu()

		if choice == 3 {
			break
		}

		switch choice {
			case 1:
				newTask := addTask()
				tasks = append(tasks, newTask)
				idTaskMap[newTask.Id] = &tasks[len(tasks) - 1]
			case 2:
				markTaskAsDone(idTaskMap)
		}
	}

	
}

func menu() int {

	retry:
	fmt.Println("Menu:")
	fmt.Println("1. Add Task")
	fmt.Println("2. Mark Task as Done")
	fmt.Println("3. Exit")
	reader := bufio.NewReader(os.Stdin)
	choiceStr, _ := reader.ReadString('\n')
	choice,_ := strconv.Atoi(strings.TrimSpace(choiceStr));
	
	if choice > 0 && choice < 4 {
		return choice
	} else {
		fmt.Println("Invalid Choice!")
		goto retry
	}
}

func addTask() Task {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter Task ID (numeric)");
	idStr, _ := reader.ReadString('\n')
	id, _ := strconv.Atoi(strings.TrimSpace(idStr))
	
	fmt.Println("Enter Task")
	taskRawStr,_ := reader.ReadString('\n')
	task := strings.TrimSpace(taskRawStr)

	return Task{
		Id: id,
		Task: task,
		IsDone: false,
	}
}

func markTaskAsDone(idTaskMap map[int]*Task) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter ID")
	idStr, _ := reader.ReadString('\n')
	id, _ := strconv.Atoi(strings.TrimSpace(idStr))

	task, err := idTaskMap[id]

	fmt.Println(err)
	if !err {
		fmt.Println("Invalid ID")
	} else {
		task.IsDone = true;
	}
}

func printSummary(idTaskMap map[int]*Task) {
	completedTasks := 0
	pendingTasks := 0

	for _, task := range idTaskMap {
		if task.IsDone {
			completedTasks++;
		} else {
			pendingTasks++;
		}
	}
	
	fmt.Println("Total Tasks", len(idTaskMap))
	fmt.Println("Completed Tasks", completedTasks)
	fmt.Println("Pending Tasks", pendingTasks)
}