package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func fibonacci() func() int {
	first, second := 0, 1
	return func() int {
		ret := first
		first, second = second, first+second
		return ret
	}
}

type FibMes struct {
	 Fibonacci int
     Index  int
}

func fail(id, fib int) {

	fmsg := &FibMes{Fibonacci: fib, Index: id}
	msgJson, err := json.Marshal(fmsg)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Correct answer:")
	fmt.Println(string(msgJson))

}

func getInput(input chan int) {
	for {

		in := bufio.NewReader(os.Stdin)
		inpt, _ := in.ReadString('\n')
		inpt = strings.Replace(inpt, "\n", "", -1)
		result, err := strconv.Atoi(inpt)
		if err != nil {
			fmt.Println("Please input only numbers.")
			input <- -1
		} else {

			input <- result
		}
	}
}

func main() {
	fib := fibonacci()
	userErrors := 0
	userCorrects := 0
	fmt.Println("Hey there! Let's play Fibonacci game.\nYou must input a number in not more than 10 secs. You have 3 lives...\n(Hint: sequence starts from 0)\nGO!")

	input := make(chan int, 1)
	go getInput(input)

  for i := 0; (userCorrects < 10 && userErrors < 3); i++ {
		fib_cur := fib()
		fmt.Println("Input Fibonacci number:")

		select {
		case inpt := <-input:
			if inpt == fib_cur {
				fmt.Println("Yay, you're right!")
				userCorrects++
			} else if inpt != fib_cur {
				fmt.Println("You're wrong.")
				fail(i, fib_cur)
				userErrors++
			}
		case <-time.After(10 * time.Second):
			fmt.Println("Your 10 secs expired.")
			fail(i, fib_cur)
			userErrors++

		}
	}

	if userErrors > 2 {
		fmt.Println("You lost :(")
	} else {
		fmt.Println("You won :)")
	}
}
