package main

import (
	"fmt"
	"os"
	"strings"
)

type CoffeeMachine struct {
	water, milk, beans, cups, money int
}

func displayStock(cm *CoffeeMachine) {
	fmt.Println("The coffee machine has:")
	fmt.Printf("%d ml of  water\n", cm.water)
	fmt.Printf("%d ml of milk\n", cm.milk)
	fmt.Printf("%d g of coffee beans\n", cm.beans)
	fmt.Printf("%d disposable cups\n", cm.cups)
	fmt.Printf("$%d of money\n", cm.money)
	fmt.Println()
}

func main() {
	cm := &CoffeeMachine{400, 540, 120, 9, 550}

	for {
		fmt.Println("Write action (buy, fill, take, remaining, exit):")
		var action string
		fmt.Scanln(&action)

		switch strings.ToLower(action) {
		case "buy":
			buyCoffee(cm)
		case "fill":
			fillInv(cm)
		case "take":
			emptyRegister(cm)
		case "remaining":
			displayStock(cm)
		case "exit":
			os.Exit(0)
		}
	}
}

func emptyRegister(cm *CoffeeMachine) {
	fmt.Printf("I gave you $%d\n", cm.money)
	cm.money = 0
	fmt.Println()
}

func fillInv(cm *CoffeeMachine) {
	fmt.Println("Write how many ml of water you want to add:")
	var addWater int
	fmt.Scanln(&addWater)
	fmt.Println("Write how many ml of milk you want to add:")
	var addMilk int
	fmt.Scanln(&addMilk)
	fmt.Println("Write how many grams of coffee beans you want to add:")
	var addBeans int
	fmt.Scanln(&addBeans)
	fmt.Println("Write how many grams of disposable cups you want to add:")
	var addCups int
	fmt.Scanln(&addCups)

	cm.water += addWater
	cm.milk += addMilk
	cm.beans += addBeans
	cm.cups += addCups

	fmt.Println()

}

func makeEspresso(cm *CoffeeMachine) {
	cm.water -= 250
	cm.beans -= 16
	cm.cups -= 1
	cm.money += 4
	fmt.Println("I have enough resources, making you a coffee!")
	fmt.Println()
}

func makeLatte(cm *CoffeeMachine) {
	cm.water -= 350
	cm.milk -= 75
	cm.beans -= 20
	cm.cups -= 1
	cm.money += 7
	fmt.Println("I have enough resources, making you a coffee!")
	fmt.Println()
}

func makeCappuccino(cm *CoffeeMachine) {
	cm.water -= 200
	cm.milk -= 100
	cm.beans -= 12
	cm.cups -= 1
	cm.money += 6
	fmt.Println("I have enough resources, making you a coffee!")
	fmt.Println()
}

func buyCoffee(cm *CoffeeMachine) {
	fmt.Println("What do you want to buy? 1 - espresso, 2 - latte, 3 - cappuccino:")
	var coffeeType string
	fmt.Scanln(&coffeeType)
	switch coffeeType {
	case "1":
		if cm.water < 250 {
			fmt.Println("Sorry, not enough water")
		} else if cm.beans < 16 {
			fmt.Println("Sorry, not enough coffee beans")
		} else {
			makeEspresso(cm)
		}
	case "2":
		if cm.water < 350 {
			fmt.Println("Sorry, not enough water")
		} else if cm.milk < 75 {
			fmt.Println("Sorry, not enough milk")
		} else if cm.beans < 20 {
			fmt.Println("Sorry, not enough coffee beans")
		} else {
			makeLatte(cm)
		}
	case "3":
		if cm.water < 200 {
			fmt.Println("Sorry, not enough water")
		} else if cm.milk < 100 {
			fmt.Println("Sorry, not enough milk")
		} else if cm.beans < 12 {
			fmt.Println("Sorry, not enough coffee beans")
		} else {
			makeCappuccino(cm)
		}
	}
}

func min(a, b, c int) int {
	if a <= b && a <= c {
		return a
	} else if b <= a && b <= c {
		return b
	}
	return c
}
