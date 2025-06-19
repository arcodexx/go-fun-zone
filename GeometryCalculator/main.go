package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main()  {

	for {
		fmt.Println("Welcome to Geometry Calculator")
		shape := shapeMenu()

		if shape == 3 {
			break
		}

		switch shape {
			case 1:
				rect := Rectangle{}
				rect.SetDimensions()
				fmt.Println("Heigth - ", rect.height)
				fmt.Println("Width - ", rect.width)
				fmt.Println("Area - ", rect.Area())
				fmt.Println("Perimeter - ", rect.Perimeter())	
				
			case 2:
				square := Square{}
				square.SetDimensions()
				fmt.Println("Side - ", square.side)
				fmt.Println("Area - ", square.Area())
				fmt.Println("Perimeter - ", square.Perimeter())
		}

		

	}
	
	
}

type Rectangle struct {
	height float64
	width float64
}

func (r Rectangle) Area() float64 {
	return r.height * r.width
}

func (r Rectangle) Perimeter() float64 {
	return (2*r.height) + (2*r.width)
}

func (r *Rectangle) SetDimensions() {
	reader := bufio.NewReader(os.Stdin)

	retry:

	fmt.Println("Please enter heigth")
	heightStr, _ := reader.ReadString('\n')
	height, _ := strconv.ParseFloat(strings.TrimSpace(heightStr), 64)

	fmt.Println("Please enter width")
	widthStr, _ := reader.ReadString('\n')
	width, _ := strconv.ParseFloat(strings.TrimSpace(widthStr), 64)

	if width == height {
		fmt.Println("The given dimensions are not of rectangle, Try Again!")
		goto retry
	}

	r.width = width
	r.height = height
}

type Square struct {
	side float64
}

func (s Square) Area() float64 {
	return s.side*s.side
}

func (s Square) Perimeter() float64 {
	return 4*s.side
}

func (s *Square) SetDimensions() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Please enter side")
	sideStr, _ := reader.ReadString('\n')
	side, _ := strconv.ParseFloat(strings.TrimSpace(sideStr), 64)

	s.side = side
}

func shapeMenu() int  {
	
	retryShapeMenu:
	fmt.Println("Select Shape")
	fmt.Println("1. Rectangle")
	fmt.Println("2. Square")
	fmt.Println("Enter 3 to exit")

	reader := bufio.NewReader(os.Stdin)
	choiceStr,_ := reader.ReadString('\n')
	choice,_ := strconv.Atoi(strings.TrimSpace(choiceStr))

	if choice > 0 && choice < 4 {
		return choice
	} else {
		fmt.Println("Invalid Value")
		goto retryShapeMenu
	}
}