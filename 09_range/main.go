package main

import "fmt"

func main() {
	ids := []int{32, 44, 11, 55, 2}

	// Loop Through ids
	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	fmt.Printf("\n")

	// Not using index
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	// Add ids together
	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Printf("\n%d\n", sum)

	// Range with map
	emails := map[string]string{"Bob": "bob@gmail.com", "Sharon": "Sharon@gmail.com"}

	emails["Mike"] = "mike@gmail.com"

	for k, v := range emails {
		fmt.Printf("%s: %s\n", k, v)
	}

	fmt.Print("\n")

	for k := range emails {
		fmt.Println("Name :	" + k)
	}
}
