package main

import (
	"fmt"
	"github.com/kyokomi/emoji/v2"
	"strconv"
)

func greet(name, year string) {
	fmt.Println("Hello! My name is " + name + ".")
	fmt.Println("I was created in " + year + ".")
}

func showName() {
	var name string
	fmt.Println("Please, remind me your name.")
	fmt.Scan(&name)
	fmt.Println("What a great name you have, " + name + "!")
}

func guessAge() {
	var rem3, rem5, rem7, age int

	fmt.Println("Let me guess your age.")
	fmt.Println("Enter remainders of dividing your age by 3, 5 and 7.")
	fmt.Scan(&rem3, &rem5, &rem7)

	age = (rem3*70 + rem5*21 + rem7*15) % 105
	fmt.Println("Your age is " + strconv.Itoa(age) + "; that's a good time to start programming!")
}

func count() {
	var n int

	fmt.Println("Now I will prove to you that I can count to any number you want.")
	fmt.Scan(&n)
	for i := 0; i <= n; i++ {
		fmt.Printf("%d!\n", i)
	}
}

func startQuiz() {
	fmt.Println("Let's test your programming knowledge.")
	var choice int

	fmt.Println("Why do we choose to write code in Golang?")
	fmt.Println("1. So we can write more DevOps tools.")
	fmt.Println("2. To be one of the cool kids.")
	fmt.Println("3. Because you don't have to type semicolon at the end.")
	fmt.Println("4. Because it is pragmatic and easy to follow")

	for choice != 4 {
		fmt.Scan(&choice)
		fmt.Println("Please, try again.")
	}
}

func sayGoodbye() {
	sayBye := "Congratulations, have a nice day!"
	fmt.Println(sayBye)
	programmer := emoji.Sprint(":man_technologist:")
	emoji.Printf(":woman_technologist:")
	color.Green(" Learning about modules! " + programmer)

}

func main() {
	greet("Aid", "2020") // change it as you need
	showName()
	guessAge()
	count()
	startQuiz()
	sayGoodbye()
}
