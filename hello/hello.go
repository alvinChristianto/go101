package main

import (
	"fmt"
	"log"
	"math"

	"example.com/greetings"
)

const date int = 2024
const thisText string = "constant thisText"

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// Request a greeting message.
	// message, err := greetings.Hello("Gladys")
	// If an error was returned, print it to the console and
	// exit the program.
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// If no error was returned, print the returned message
	// to the console.

	// A slice of names.
	names := []string{"Gladys", "Samantha", "Darrin"}

	// Request greeting messages for the names.
	multipleHello, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}

	singleHello, err := greetings.Hello("Gladys")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("---------------> Hello")
	fmt.Println(singleHello)
	// If no error was returned, print the returned map of
	// messages to the console.
	fmt.Println(multipleHello)

	fmt.Println("\n---------------> values")
	fmt.Println("go" + "lang")

	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)
	fmt.Println("7.0/3.0 =", 7.12/3.45)
	fmt.Printf("type of 7.12 %T\n", 7.12)

	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)

	fmt.Println("\n---------------> variables")

	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f := "apple"
	fmt.Println(f)

	fmt.Println("\n---------------> constants")
	fmt.Println(`   In JavaScript, const only creates a "constant reference". This means you cannot point the variable to something else, but if the value is an object or array, you can still modify its contents.
   In Go, a constant is the value itself. Because Go constants are restricted to simple scalars (numbers, strings, booleans), there are no "internal parts" to modify. A Go constant is immutable in every sense of the word.`)

	fmt.Println("date:", date)
	fmt.Println("thisText:", thisText)

	const n = 500000000

	const dd = 3e20 / n
	fmt.Println(dd)

	fmt.Println(int64(dd))

	fmt.Println(math.Sin(n))

	fmt.Println("\n---------------> for loops")
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}

	for i := range 3 {
		fmt.Println("range", i)
	}

	for {
		fmt.Println("loop")
		break
	}

	for n := range 6 {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
