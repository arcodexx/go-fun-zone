package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main()  {
	type Item struct {
		ID string
		Name string
		Quantity int
	}

	inventory := []Item{
		{ID: "APL", Name: "Apple", Quantity: 10},
		{ID: "BAN", Name: "Banana", Quantity: 20},
		{ID: "CHRY", Name: "Cherry", Quantity: 15},
	}

	mapInventory := make(map[string]*Item)
	mapInventory["APL"] = &inventory[0]
	mapInventory["BAN"] = &inventory[1]
	mapInventory["CHRY"] = &inventory[2]

	fmt.Println("Inventory Items:");
	fmt.Println(inventory);

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter Item ID:");
	id, _ := reader.ReadString('\n')
	fmt.Println("New Item Name:");
	name, _ := reader.ReadString('\n')
	fmt.Println("New Item Quantity:");
	quantityStr, _ := reader.ReadString('\n')
	quantity, _ := strconv.Atoi(strings.TrimSpace(quantityStr))

	inventory = append(inventory, Item{strings.TrimSpace(id), strings.TrimSpace(name), quantity})
	mapInventory[strings.TrimSpace(id)] = &inventory[len(inventory)-1]


	fmt.Println("Inventory Items:");
	fmt.Println(inventory);

	fmt.Println("Enter Item ID to update:");
	updateId, _ := reader.ReadString('\n')
	updateId = strings.TrimSpace(updateId)

	fmt.Println("Enter New Quatity:");
	newQuantityStr, _ := reader.ReadString('\n')
	newQuantity, _ := strconv.Atoi(strings.TrimSpace(newQuantityStr))

	updateItem := mapInventory[updateId]

	updateItem.Quantity = newQuantity

	fmt.Println("Updated Inventory Items:");
	fmt.Println(inventory)

}