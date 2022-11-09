package main

import "fmt"

func main() {
	a := 10

	if a == 10 {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}

	switch a {
	case 10:
		fmt.Println(10)
	default:
		fmt.Println("Default")
	}

	const b int = 11

	fmt.Println(b)

	lastname := "Chouybumrung"

	fmt.Printf("Samut %s", lastname)

}
