package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// structure of the bill
type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// function to create a new bill
func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}
	return b
}

// function to format the bill
func (b bill) format() string {
	formatedString := "Bill breakdown: \n"
	var total float64 = 0

	// loop through the list all the items to get the total
	for k, v := range b.items {
		formatedString += fmt.Sprintf("%-10v ...$%v \n", k+":", v)
		total += v
	}

	// add the tip
	formatedString += fmt.Sprintf("%-10v ...$%v\n", "tip:", b.tip)

	// calculate the total
	formatedString += fmt.Sprintf("%-10v ...$%0.2f", "total:", total)
	return formatedString
}

// update tip
func (b *bill) updateTip(tip float64) {
	b.tip = tip
}

// add a new item to the bill
func (b bill) addItem(name string, price float64) {
	b.items[name] = price
}

// save the bill
func (b *bill) save() {
	data := []byte(b.format())
	err := os.WriteFile("bills/"+b.name+".txt", data, 0664)
	if err != nil {
		panic(err)
	}
	fmt.Println("You saved the bill - ", b.name)
}

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)

	name, _ := getInput("Create a new bill name: ", reader)

	b := newBill(name)
	fmt.Println("Created the bill - ", b.name)

	return b
}

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)
	opt, _ := getInput("Choose option (add - add item, tip - add tip, save - save bill, exit - to exit :  ", reader)

	switch opt {
	case "add":
		name, _ := getInput("Item name: ", reader)
		price, _ := getInput("Item price: ", reader)
		p, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("The price must be a number")
			promptOptions(b)
		}
		b.addItem(name, p)

		fmt.Println("Item added - ", name, price)
		promptOptions(b)

	case "tip":
		tip, _ := getInput("Enter tip amount ($): ", reader)
		t, err := strconv.ParseFloat(tip, 64)
		if err != nil {
			fmt.Println("The tip must be a number")
			promptOptions(b)
		}
		b.updateTip(t)

		fmt.Println("Tip added - ", tip)
		promptOptions(b)

	case "saves":
		b.save()
		promptOptions(b)

	case "exit":
		break

	default:
		fmt.Printf("That was not a valid option \n")
		promptOptions(b)
	}
}

func main() {
	mybill := createBill()
	promptOptions(mybill)
}
