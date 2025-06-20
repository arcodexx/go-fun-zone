package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
)


const url = "https://v6.exchangerate-api.com/v6/<YOUR-API-KEY>/"

func main()  {

	for {
		choice := showMenu()
		
		if choice == 3 {
			break
		}
		
		switch choice {
			case 1:
				printSupportedCodeData()
			case 2:
				converterInit()

		}
		
	}

	

}

func showMenu() int  {
	reader := bufio.NewReader(os.Stdin)

	rePrintMenu:

	fmt.Println("Welcome to Currency Converter")
	fmt.Println("1. Show Supported Codes")
	fmt.Println("2. Convert")
	fmt.Println("3. Exit")

	choiceStr,_ := reader.ReadString('\n')
	choice, _ := strconv.Atoi(strings.TrimSpace(choiceStr))

	if choice > 0 && choice < 4 {
		return choice
	} else {
		fmt.Println("Invalid Choice")
		goto rePrintMenu
	}
}

func printSupportedCodeData() {

	type responseObj struct {
		Result string `json:"result"`
		Docs string `json:"documentation"`
		Terms string `json:"terms_of_use"` 
		Data [][2]string `json:"supported_codes"`
	}

	var countryCodes responseObj

	response, err := http.Get(url+"codes")
	checkErrPanic(err)
	defer response.Body.Close()

	responseByte, _ := io.ReadAll(response.Body)

	json.Unmarshal(responseByte, &countryCodes)

	for index, codeData := range countryCodes.Data {
		fmt.Printf("%v. %v - %v \n", index, codeData[0], codeData[1])
	}
}

func converterInit() {
	
	type Response struct {
		Result string `json:"result"`
		ConversionRate float64 `json:"conversion_rate"`
		ConvertedValue float64 `json:"conversion_result"`
	}
	
	
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter Base Country Code")
	base, _ := reader.ReadString('\n')
	base = strings.TrimSpace(base)
	
	fmt.Println("Enter Target Country Code")
	target, _ := reader.ReadString('\n')
	target = strings.TrimSpace(target)

	fmt.Println("Enter Amount")
	amount, _ := reader.ReadString('\n')
	amount = strings.TrimSpace(amount)

	rawResponse, err := http.Get(url+"/pair/"+base+"/"+target+"/"+amount)
	checkErrPanic(err)
	defer rawResponse.Body.Close()

	var response Response

	responseByte, _ := io.ReadAll(rawResponse.Body)
	
	json.Unmarshal(responseByte, &response)

	fmt.Printf("Conversion Rate from %v to %v is %v \n", base, target, response.ConversionRate)
	fmt.Printf("%v %v will be converted to %v %v \n", amount, base, response.ConvertedValue, target)
}


func checkErrPanic(err error) {
	if err != nil {
		panic(err)
	}
}